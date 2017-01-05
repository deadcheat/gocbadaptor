package gocbadaptor

import (
	"log"

	"github.com/couchbase/gocb"
	"github.com/deadcheat/gocbadaptor/conf"
)

// DefaultCouchAdaptor default adaptor struct
type DefaultCouchAdaptor struct {
	Environment *conf.Env
	CouchBucket *gocb.Bucket
}

// NewDefaultCouchAdaptor 新しいAdaptorインスタンスを作成
func NewDefaultCouchAdaptor() *DefaultCouchAdaptor {
	return &DefaultCouchAdaptor{}
}

// Env return Environment
func (adaptor *DefaultCouchAdaptor) Env() *conf.Env {
	return adaptor.Environment
}

// Bucket return CouchBucket
func (adaptor *DefaultCouchAdaptor) Bucket() *gocb.Bucket {
	return adaptor.CouchBucket
}

// Open open adaptor.CouchBucket using arguments
func (adaptor *DefaultCouchAdaptor) Open(connection, bucket, password string, expiry uint32) (err error) {
	adaptor.Environment = &conf.Env{
		ConnectString: connection,
		BucketName:    bucket,
		Password:      password,
		CacheExpiry:   expiry,
	}
	e := *adaptor.Environment
	adaptor.CouchBucket, err = e.OpenBucket()
	return
}

// OpenWithConfig open adaptor.CouchBucket using config struct
func (adaptor *DefaultCouchAdaptor) OpenWithConfig(env *conf.Env) (err error) {
	adaptor.Environment = env
	e := *env
	adaptor.CouchBucket, err = e.OpenBucket()
	return
}

// Get invoke gocb.adaptor.Bucket.Get
func (adaptor *DefaultCouchAdaptor) Get(key string) (cas gocb.Cas, data []byte, ok bool) {
	if adaptor.CouchBucket == nil {
		log.Printf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil, false
	}
	b := *adaptor.CouchBucket
	cas, err := b.Get(key, &data)
	if err != nil {
		log.Printf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, nil, false
	}
	log.Printf("hit key: %s", key)
	return cas, data, true
}

// Insert invoke gocb.adaptor.Bucket.Insert
func (adaptor *DefaultCouchAdaptor) Insert(key string, data []byte) (cas gocb.Cas, ok bool) {
	if adaptor.CouchBucket == nil {
		return 0, false
	}
	b := *adaptor.CouchBucket
	cas, err := b.Insert(key, data, adaptor.Environment.CacheExpiry)
	if err != nil {
		log.Printf("Couldn't insert for key: %s or err: %+v \n", key, err)
		return cas, false
	}
	log.Printf("insert to adaptor.CouchBucket key: %s", key)
	return cas, true
}

// Upsert invoke gocb.adaptor.Bucket.Upsert
func (adaptor *DefaultCouchAdaptor) Upsert(key string, data []byte) (cas gocb.Cas, ok bool) {
	b := *adaptor.CouchBucket
	cas, err := b.Upsert(key, data, adaptor.Environment.CacheExpiry)
	if err != nil {
		log.Printf("Couldn't upsert for key: %s or err: %+v \n", key, err)
		return cas, false
	}
	log.Printf("upsert to adaptor.CouchBucket key: %s", key)
	return cas, true
}

// N1qlQuery prepare query and execute
func (adaptor *DefaultCouchAdaptor) N1qlQuery(q string, params interface{}) (r gocb.QueryResults, err error) {
	if adaptor.CouchBucket == nil {
		return nil, nil
	}
	nq := gocb.NewN1qlQuery(q)
	b := *adaptor.CouchBucket
	r, err = b.ExecuteN1qlQuery(nq, params)
	if err != nil {
		log.Printf("Couldn't execute query for query: %s params: %+v or err: %+v \n", q, params, err)
		return r, err
	}
	log.Printf("succeeded to execute query: %s , params: %+v", q, params)
	return r, err
}

// ExpiresIn overwrite Env.CacheExpiry
func (adaptor *DefaultCouchAdaptor) ExpiresIn(sec uint32) {
	adaptor.Environment.CacheExpiry = sec
}

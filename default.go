package gocbadaptor

import (
	"log"

	"github.com/couchbase/gocb"
	"github.com/deadcheat/gocbadaptor/conf"
)

// DefaultCouchAdaptor default adaptor struct
type DefaultCouchAdaptor struct {
	Env    *conf.Env
	Bucket *gocb.Bucket
}

// NewDefaultCouchAdaptor 新しいAdaptorインスタンスを作成
func NewDefaultCouchAdaptor() *DefaultCouchAdaptor {
	return &DefaultCouchAdaptor{}
}

// Open open adaptor.Bucket using arguments
func (adaptor DefaultCouchAdaptor) Open(connection, bucket, password *string, expiry uint32) CouchBaseAdaptor {
	adaptor.Env = &conf.Env{
		ConnectString: connection,
		BucketName:    bucket,
		Password:      password,
		CacheExpiry:   expiry,
	}
	e := *adaptor.Env
	adaptor.Bucket = e.OpenBucket()
	return adaptor
}

// OpenWithConfig open adaptor.Bucket using config struct
func (adaptor DefaultCouchAdaptor) OpenWithConfig(env *conf.Env) CouchBaseAdaptor {
	adaptor.Env = env
	e := *env
	adaptor.Bucket = e.OpenBucket()
	return adaptor
}

// Get invoke gocb.adaptor.Bucket.Get
func (adaptor DefaultCouchAdaptor) Get(key string) (cas gocb.Cas, data []byte, ok bool) {
	if adaptor.Bucket == nil {
		log.Printf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil, false
	}
	b := *adaptor.Bucket
	cas, err := b.Get(key, &data)
	if err != nil {
		log.Printf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, nil, false
	}
	log.Printf("hit key: %s", key)
	return cas, data, true
}

// Insert invoke gocb.adaptor.Bucket.Insert
func (adaptor DefaultCouchAdaptor) Insert(key string, data []byte) (cas gocb.Cas, ok bool) {
	if adaptor.Bucket == nil {
		return 0, false
	}
	b := *adaptor.Bucket
	cas, err := b.Insert(key, data, adaptor.Env.CacheExpiry)
	if err != nil {
		log.Printf("Couldn't insert for key: %s or err: %+v \n", key, err)
		return cas, false
	}
	log.Printf("insert to adaptor.Bucket key: %s", key)
	return cas, true
}

// Upsert invoke gocb.adaptor.Bucket.Upsert
func (adaptor DefaultCouchAdaptor) Upsert(key string, data []byte) (cas gocb.Cas, ok bool) {
	if adaptor.Bucket == nil {
		return 0, false
	}
	b := *adaptor.Bucket
	cas, err := b.Upsert(key, data, adaptor.Env.CacheExpiry)
	if err != nil {
		log.Printf("Couldn't upsert for key: %s or err: %+v \n", key, err)
		return cas, false
	}
	log.Printf("upsert to adaptor.Bucket key: %s", key)
	return cas, true
}

// N1qlQuery prepare query and execute
func (adaptor DefaultCouchAdaptor) N1qlQuery(q string, params interface{}) (r gocb.QueryResults, err error) {
	if adaptor.Bucket == nil {
		return nil, nil
	}
	nq := gocb.NewN1qlQuery(q)
	b := *adaptor.Bucket
	r, err = b.ExecuteN1qlQuery(nq, params)
	if err != nil {
		log.Printf("Couldn't execute query for query: %s params: %+v or err: %+v \n", q, params, err)
		return r, err
	}
	r.Close()
	log.Printf("succeeded to execute query: %s , params: %+v", q, params)
	return r, err
}

// ExpiresIn overwrite Env.CacheExpiry
func (adaptor DefaultCouchAdaptor) ExpiresIn(sec uint32) CouchBaseAdaptor {
	adaptor.Env.CacheExpiry = sec
	return adaptor
}

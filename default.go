package gocbadaptor

import (
	"errors"
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
func (a *DefaultCouchAdaptor) Env() *conf.Env {
	return a.Environment
}

// Bucket return CouchBucket
func (a *DefaultCouchAdaptor) Bucket() *gocb.Bucket {
	return a.CouchBucket
}

// Open open a.CouchBucket using arguments
func (a *DefaultCouchAdaptor) Open(connection, bucket, password string, expiry uint32) (err error) {
	if a == nil {
		return
	}
	a.Environment = &conf.Env{
		ConnectString: connection,
		BucketName:    bucket,
		Password:      password,
		CacheExpiry:   expiry,
	}
	e := *a.Environment
	a.CouchBucket, err = e.OpenBucket()
	return
}

// OpenWithConfig open a.CouchBucket using config struct
func (a *DefaultCouchAdaptor) OpenWithConfig(env *conf.Env) (err error) {
	if a == nil {
		return
	}
	a.Environment = env
	e := *env
	a.CouchBucket, err = e.OpenBucket()
	return
}

// Get invoke gocb.a.Bucket.Get
func (a *DefaultCouchAdaptor) Get(key string) (cas gocb.Cas, data []byte, err error) {
	if a == nil || a.CouchBucket == nil {
		log.Printf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil, nil
	}
	b := *a.CouchBucket
	cas, err = b.Get(key, &data)
	if err != nil {
		log.Printf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, nil, err
	}
	log.Printf("hit key: %s", key)
	return cas, data, nil
}

// Insert invoke gocb.a.Bucket.Insert
func (a *DefaultCouchAdaptor) Insert(k string, d []byte) (gocb.Cas, error) {
	return a.update(insert, k, d)
}

// Upsert invoke gocb.a.Bucket.Upsert
func (a *DefaultCouchAdaptor) Upsert(k string, d []byte) (gocb.Cas, error) {
	return a.update(upsert, k, d)
}

type updateMode int

const (
	insert updateMode = iota
	upsert
)

func (a *DefaultCouchAdaptor) update(mode updateMode, key string, data []byte) (c gocb.Cas, e error) {
	if a == nil || a.CouchBucket == nil {
		return 0, nil
	}
	b := *a.CouchBucket
	if mode == insert {
		c, e = b.Insert(key, data, a.Environment.CacheExpiry)
	} else if mode == upsert {
		c, e = b.Upsert(key, data, a.Environment.CacheExpiry)
	} else {
		log.Fatal(errors.New("update should not call insert or upsert mode"))
	}
	if e != nil {
		log.Printf("Couldn't send data for key: %s or err: %+v \n", key, e)
		return c, e
	}
	log.Printf("sent data to a.CouchBucket key: %s", key)
	return c, nil
}

// N1qlQuery prepare query and execute
func (a *DefaultCouchAdaptor) N1qlQuery(q string, params interface{}) (r gocb.QueryResults, err error) {
	if a == nil || a.CouchBucket == nil {
		return nil, nil
	}
	nq := gocb.NewN1qlQuery(q)
	b := *a.CouchBucket
	r, err = b.ExecuteN1qlQuery(nq, params)
	if err != nil {
		log.Printf("Couldn't execute query for query: %s params: %+v or err: %+v \n", q, params, err)
		return r, err
	}
	log.Printf("succeeded to execute query: %s , params: %+v", q, params)
	return r, err
}

// ExpiresIn overwrite Env.CacheExpiry
func (a *DefaultCouchAdaptor) ExpiresIn(sec uint32) {
	a.Environment.CacheExpiry = sec
}

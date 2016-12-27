package gocbadaptor

import (
	"log"

	"github.com/couchbase/gocb"
	"github.com/deadcheat/gocbadaptor/conf"
)

// DefaultCouchAdapter default adapter struct
type DefaultCouchAdaptor struct{}

// NewCouchAdapter 新しいAdapterインスタンスを作成
func NewCouchAdapter() *DefaultCouchAdaptor {
	return &DefaultCouchAdapter{}
}

// Open open bucket using config struct
func (*DefaultCouchAdapter) Open(c *conf.Env) *gocb.Bucket {
	return c.OpenBucket()
}

// Get invoke gocb.Bucket.Get
func (*DefaultCouchAdapter) Get(b *gocb.Bucket, key string) (cas gocb.Cas, data []byte, ok bool) {
	if b == nil {
		log.Printf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil, false
	}
	cas, err := b.Get(key, &data)
	if err != nil {
		log.Printf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, nil, false
	}
	log.Printf("hit key: %s", key)
	return cas, data, true
}

// Insert invoke gocb.Bucket.Insert
func (*DefaultCouchAdapter) Insert(b *gocb.Bucket, key string, data []byte, expiry uint32) (cas gocb.Cas, ok bool) {
	if b == nil {
		return 0, false
	}
	cas, err := b.Insert(key, data, expiry)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't insert for key: %s or err: %+v \n", key, err)
		return cas, false
	}
	log.Printf("insert to bucket key: %s", key)
	return cas, true
}

// Upsert invoke gocb.Bucket.Upsert
func Upsert(b *gocb.Bucket, key string, data []byte, expiry uint32) (cas gocb.Cas, ok bool) {
	if b == nil {
		return 0, false
	}
	cas, err := b.Upsert(key, data, expiry)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't upsert for key: %s or err: %+v \n", key, err)
		return cas, false
	}
	log.Printf("insert to bucket key: %s", key)
	return cas, true
}

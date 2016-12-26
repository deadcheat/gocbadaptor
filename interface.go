package gocbadaptor

import (
	"github.com/deadcheat/gocbadaptor/conf"

	"github.com/couchbase/gocb"
)

// CouchAdapter CouchBase接続アダプタ
type CouchAdapter interface {
	Open(couchenv *conf.Env) *gocb.Bucket
	Get(b *gocb.Bucket, key string) (cas gocb.Cas, data []byte, ok bool)
	Insert(b *gocb.Bucket, key string, data []byte, expiry uint32) (cas gocb.Cas, ok bool)
	Upsert(b *gocb.Bucket, key string, data []byte, expiry uint32) (cas gocb.Cas, ok bool)
}

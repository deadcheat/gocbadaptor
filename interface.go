package gocbadaptor

import (
	"github.com/deadcheat/gocbadaptor/conf"

	"github.com/couchbase/gocb"
)

// CouchBaseAdaptor CouchBase接続アダプタ
type CouchBaseAdaptor interface {
	Open(connection, bucket, password string, expiry uint32) CouchBaseAdaptor
	OpenWithConfig(env *conf.Env) CouchBaseAdaptor
	Env() *conf.Env
	Bucket() *gocb.Bucket
	ExpiresIn(sec uint32) CouchBaseAdaptor
	Get(key string) (cas gocb.Cas, data []byte, ok bool)
	Insert(key string, data []byte) (cas gocb.Cas, ok bool)
	Upsert(key string, data []byte) (cas gocb.Cas, ok bool)
	N1qlQuery(q string, params interface{}) (r gocb.QueryResults, err error)
}

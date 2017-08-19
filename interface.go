package gocbadaptor

import (
	"github.com/deadcheat/gocbadaptor/conf"

	"github.com/couchbase/gocb"
)

// CouchBaseAdaptor CouchBase connect adaptor
type CouchBaseAdaptor interface {
	Bucket() *gocb.Bucket
	Env() *conf.Env
	ExpiresIn(sec uint32)
	Get(key string) (cas gocb.Cas, data []byte, err error)
	Insert(key string, data []byte) (cas gocb.Cas, err error)
	N1qlQuery(q string, params interface{}) (r gocb.QueryResults, err error)
	N1qlQueryWithMode(mode *gocb.ConsistencyMode, q string, params interface{}) (r gocb.QueryResults, err error)
	Open(connection, bucket, password string, expiry uint32) (err error)
	OpenWithConfig(env *conf.Env) (err error)
	Upsert(key string, data []byte) (cas gocb.Cas, err error)
}

// Loggerable logging interface
type Loggerable interface {
	Log(...interface{})
	Logf(format string, v ...interface{})
}

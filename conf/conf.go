package conf

import "github.com/couchbase/gocb"

// Env struct for configure connection
type Env struct {
	ConnectString string `mapstructure:"connection"`
	BucketName    string `mapstructure:"bucketname"`
	Password      string `mapstructure:"password"`
	CacheExpiry   uint32 `mapstructure:"expiry"`
}

// OpenBucket CouchBaseへの接続開始
func (env *Env) OpenBucket() (bucket *gocb.Bucket, err error) {
	cluster, err := gocb.Connect(env.ConnectString)
	if err != nil {
		return nil, err
	}
	bucket, err = cluster.OpenBucket(env.BucketName, env.Password)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

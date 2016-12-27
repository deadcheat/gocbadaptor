package conf

import (
	"fmt"
	"log"

	"github.com/couchbase/gocb"
)

// Env struct for configure connection
type Env struct {
	ConnectString string `mapstructure:"connection"`
	BucketName    string `mapstructure:"bucketname"`
	Password      string `mapstructure:"password"`
	CacheExpiry   uint32 `mapstructure:"expiry"`
}

// OpenBucket CouchBaseへの接続開始
func (env *Env) OpenBucket() *gocb.Bucket {
	cluster, err := gocb.Connect(env.ConnectString)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to open cluster: %s", env.ConnectString), err)
		return nil
	}
	log.Printf("Succeeded to open Cluster: %s", env.ConnectString)

	bucket, err := cluster.OpenBucket(env.BucketName, env.Password)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to open bucket: [%s](%s)", env.BucketName, env.Password), err)
		return nil
	}
	log.Printf("Succeeded to open Bucket %s:%s", env.ConnectString, env.BucketName)
	return bucket
}

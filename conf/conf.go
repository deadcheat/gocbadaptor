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
func (c *Env) OpenBucket() *gocb.Bucket {
	cluster, err := gocb.Connect(c.ConnectString)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to open cluster: %s", c.ConnectString), err)
		return nil
	}
	log.Printf("Succeeded to open Cluster: %s", c.ConnectString)

	bucket, err := cluster.OpenBucket(c.BucketName, c.Password)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to open bucket: [%s](%s)", c.BucketName, c.Password), err)
		return nil
	}
	log.Printf("Succeeded to open Bucket %s:%s", c.ConnectString, c.BucketName)
	return bucket
}

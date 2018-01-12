package conf

import "github.com/couchbase/gocb"

// Env struct for configure connection
type Env struct {
	ClusterUser     string `mapstructure:"clusteruser"`
	ClusterPassword string `mapstructure:"clusterpassword"`
	ConnectString   string `mapstructure:"connection"`
	BucketName      string `mapstructure:"bucketname"`
	Password        string `mapstructure:"password"`
	CacheExpiry     uint32 `mapstructure:"expiry"`
}

// OpenBucket CouchBaseへの接続開始
func (env *Env) OpenBucket() (bucket *gocb.Bucket, err error) {
	var cluster *gocb.Cluster
	cluster, err = gocb.Connect(env.ConnectString)
	if err == gocb.ErrAuthError {
		err = cluster.Authenticate(gocb.PasswordAuthenticator{
			Username: env.ClusterUser,
			Password: env.ClusterPassword,
		})
	}
	if err != nil {
		return nil, err
	}
	bucket, err = cluster.OpenBucket(env.BucketName, env.Password)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

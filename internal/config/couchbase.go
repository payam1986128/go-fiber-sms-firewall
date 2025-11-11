package config

import (
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
)

type CouchbaseConfig struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
}

func InitCouchbase() (*CouchbaseConfig, error) {
	uri := os.Getenv("COUCHBASE_URI")
	user := os.Getenv("COUCHBASE_USERNAME")
	pass := os.Getenv("COUCHBASE_PASSWORD")
	bucketName := os.Getenv("COUCHBASE_BUCKET")

	if uri == "" || user == "" || pass == "" || bucketName == "" {
		return nil, fmt.Errorf("one of COUCHBASE_* env vars missing")
	}

	cluster, err := gocb.Connect(uri, gocb.ClusterOptions{Username: user, Password: pass})
	if err != nil {
		return nil, err
	}

	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		return nil, err
	}

	config := &CouchbaseConfig{
		Cluster: cluster,
		Bucket:  bucket,
	}

	defer CloseCouchbase(config)

	return config, nil
}

func CloseCouchbase(config *CouchbaseConfig) {
	if config != nil {
		_ = config.Cluster.Close(nil)
	}
}

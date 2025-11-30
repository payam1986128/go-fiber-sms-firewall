package config

import (
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
)

func InitCouchbase() (*gocb.Cluster, *gocb.Bucket, error) {
	uri := os.Getenv("COUCHBASE_URI")
	user := os.Getenv("COUCHBASE_USERNAME")
	pass := os.Getenv("COUCHBASE_PASSWORD")
	bucketName := os.Getenv("COUCHBASE_BUCKET")

	if uri == "" || user == "" || pass == "" || bucketName == "" {
		return nil, nil, fmt.Errorf("one of COUCHBASE_* env vars missing")
	}

	cluster, err := gocb.Connect(uri, gocb.ClusterOptions{Username: user, Password: pass})
	if err != nil {
		return nil, nil, err
	}

	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		return nil, nil, err
	}

	defer CloseCouchbase(cluster)

	return cluster, bucket, nil
}

func CloseCouchbase(cluster *gocb.Cluster) {
	if cluster != nil {
		_ = cluster.Close(nil)
	}
}

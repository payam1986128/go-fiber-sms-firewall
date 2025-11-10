package db

import (
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
)

var (
	Cluster    *gocb.Cluster
	Bucket     *gocb.Bucket
)

func InitCouchbase() error {
	uri := os.Getenv("COUCHBASE_URI")
	user := os.Getenv("COUCHBASE_USERNAME")
	pass := os.Getenv("COUCHBASE_PASSWORD")
	bucketName := os.Getenv("COUCHBASE_BUCKET")

	if uri == "" || user == "" || pass == "" || bucketName == "" {
		return fmt.Errorf("one of COUCHBASE_* env vars missing")
	}

	var err error
	Cluster, err = gocb.Connect(uri, gocb.ClusterOptions{Username: user, Password: pass})
	if err != nil {
		return err
	}

	Bucket = Cluster.Bucket(bucketName)
	err = Bucket.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		return err
	}

	return nil
}

func CloseCouchbase() {
	if Cluster != nil {
		_ = Cluster.Close(nil)
	}
}

package repository

import (
	"github.com/couchbase/gocb/v2"
)

var (
	limiterConditionCollection     = "limiter-conditions"
	smsCollection                  = "sms"
	suspiciousWordsCollection      = "suspicious-words"
	suspiciousCategoriesCollection = "suspicious-categories"
	userCollection                 = "users"
)

func countByQuery(cluster *gocb.Cluster, query string) (int, error) {
	data, err := cluster.Query(query, nil)
	if err != nil {
		return 0, err
	}
	var count int
	for data.Next() {
		if err := data.Row(&count); err != nil {
			return 0, err
		}
	}
	if err := data.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

package repository

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
)

var (
	limiterConditionCollection     = "limiter-conditions"
	smsCollection                  = "sms"
	suspiciousWordsCollection      = "suspicious-words"
	suspiciousCategoriesCollection = "suspicious-categories"
	userCollection                 = "users"
)

func countByQuery(query string) (int, error) {
	data, err := db.Cluster.Query(query, nil)
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

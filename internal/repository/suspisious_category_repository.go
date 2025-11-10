package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/util"
	"time"
)

func FindSuspiciousCategoriesByIds(ids []uuid.UUID) ([]entity.SuspiciousCategory, error) {
	return FindSuspiciousCategoriesByQuery(
		fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` WHERE meta().id IN [%s]", db.Bucket.Name(),
			limiterConditionCollection, util.JoinQuotedUUIDs(ids, ",")))
}

func FindSuspiciousCategoriesByQuery(query string) ([]entity.SuspiciousCategory, error) {
	data, err := db.Cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var result []entity.SuspiciousCategory
	for data.Next() {
		var record entity.SuspiciousCategory
		if err := data.Row(&record); err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func AddSuspiciousCategory(suspiciousCategory *entity.SuspiciousCategory) (uuid.UUID, error) {
	suspiciousCategory.ID = uuid.New()
	suspiciousCategory.DateTime = time.Now()
	_, err := db.Bucket.Collection(suspiciousCategoriesCollection).Insert(suspiciousCategory.ID.String(), suspiciousCategory, nil)
	return suspiciousCategory.ID, err
}

func EditSuspiciousCategory(id uuid.UUID, suspiciousCategory *entity.SuspiciousCategory) error {
	suspiciousCategory.ID = id
	_, err := db.Bucket.Collection(suspiciousCategoriesCollection).Upsert(id.String(), suspiciousCategory, nil)
	return err
}

func DeleteSuspiciousCategory(id uuid.UUID) error {
	_, err := db.Bucket.Collection(suspiciousCategoriesCollection).Remove(id.String(), nil)
	return err
}

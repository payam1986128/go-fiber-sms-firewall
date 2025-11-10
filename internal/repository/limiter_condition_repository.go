package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

func FindActiveLimiterConditions() ([]entity.LimiterCondition, error) {
	return FindLimiterConditionsByQuery(fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` WHERE active = true order by priority", db.Bucket.Name(), limiterConditionCollection))
}

func FindLimiterConditionsByQuery(query string) ([]entity.LimiterCondition, error) {
	data, err := db.Cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var result []entity.LimiterCondition
	for data.Next() {
		var record entity.LimiterCondition
		if err := data.Row(&record); err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func GetLimiterConditionByID(id uuid.UUID) (*entity.LimiterCondition, error) {
	get, err := db.Bucket.Collection(limiterConditionCollection).Get(id.String(), nil)
	if err != nil {
		return nil, err
	}
	var result entity.LimiterCondition
	if err := get.Content(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func AddLimiterCondition(limiterCondition *entity.LimiterCondition) (uuid.UUID, error) {
	limiterCondition.ID = uuid.New()
	limiterCondition.CreatedTime = time.Now()
	limiterCondition.Active = true
	_, err := db.Bucket.Collection(limiterConditionCollection).Insert(limiterCondition.ID.String(), limiterCondition, nil)
	return limiterCondition.ID, err
}

func EditLimiterCondition(id uuid.UUID, limiterCondition *entity.LimiterCondition) error {
	limiterCondition.ID = id
	_, err := db.Bucket.Collection(limiterConditionCollection).Upsert(id.String(), limiterCondition, nil)
	return err
}

func DeleteLimiterCondition(id uuid.UUID) error {
	_, err := db.Bucket.Collection(limiterConditionCollection).Remove(id.String(), nil)
	return err
}

package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/config"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

type LimiterConditionRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewLimiterConditionRepository(config *config.CouchbaseConfig) *LimiterConditionRepository {
	return &LimiterConditionRepository{
		cluster:    config.Cluster,
		collection: config.Bucket.Collection(limiterConditionCollection),
	}
}

func (repo *LimiterConditionRepository) FindActiveLimiterConditions() ([]entity.LimiterCondition, error) {
	return repo.FindLimiterConditionsByQuery(
		fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` WHERE active = true order by priority", repo.collection.Name(), limiterConditionCollection))
}

func (repo *LimiterConditionRepository) FindLimiterConditionsByQuery(query string) ([]entity.LimiterCondition, error) {
	data, err := repo.cluster.Query(query, nil)
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

func (repo *LimiterConditionRepository) GetLimiterConditionByID(id uuid.UUID) (*entity.LimiterCondition, error) {
	get, err := repo.collection.Get(id.String(), nil)
	if err != nil {
		return nil, err
	}
	var result entity.LimiterCondition
	if err := get.Content(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *LimiterConditionRepository) AddLimiterCondition(limiterCondition *entity.LimiterCondition) (uuid.UUID, error) {
	limiterCondition.ID = uuid.New()
	limiterCondition.CreatedTime = time.Now()
	limiterCondition.Active = true
	_, err := repo.collection.Insert(limiterCondition.ID.String(), limiterCondition, nil)
	return limiterCondition.ID, err
}

func (repo *LimiterConditionRepository) EditLimiterCondition(id uuid.UUID, limiterCondition *entity.LimiterCondition) error {
	limiterCondition.ID = id
	_, err := repo.collection.Upsert(id.String(), limiterCondition, nil)
	return err
}

func (repo *LimiterConditionRepository) DeleteLimiterCondition(id uuid.UUID) error {
	_, err := repo.collection.Remove(id.String(), nil)
	return err
}

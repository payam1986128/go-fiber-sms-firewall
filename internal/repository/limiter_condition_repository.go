package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/util"
	"time"
)

type LimiterConditionRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewLimiterConditionRepository(cluster *gocb.Cluster, bucket *gocb.Bucket) *LimiterConditionRepository {
	return &LimiterConditionRepository{
		cluster:    cluster,
		collection: bucket.Collection(limiterConditionCollection),
	}
}

func (repo *LimiterConditionRepository) FindActiveLimiterConditions() ([]entity.LimiterCondition, error) {
	return repo.FindAllByQuery("WHERE active = true order by priority")
}

func (repo *LimiterConditionRepository) FindAllByIds(ids []uuid.UUID) ([]entity.LimiterCondition, error) {
	return repo.FindAllByQuery(
		fmt.Sprintf("WHERE meta().id IN [%s]", util.JoinQuotedUUIDs(ids, ",")))
}

func (repo *LimiterConditionRepository) CountByQuery(whereClause string) (int, error) {
	query := fmt.Sprintf("SELECT count(meta().id) FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), limiterConditionCollection, whereClause)
	return countByQuery(repo.cluster, query)
}

func (repo *LimiterConditionRepository) FindAllByQuery(whereClause string) ([]entity.LimiterCondition, error) {
	query := fmt.Sprintf(fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), limiterConditionCollection, whereClause))
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

func (repo *LimiterConditionRepository) GetByID(id uuid.UUID) (*entity.LimiterCondition, error) {
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

func (repo *LimiterConditionRepository) Insert(limiterCondition *entity.LimiterCondition) (uuid.UUID, error) {
	limiterCondition.ID = uuid.New()
	limiterCondition.CreatedTime = time.Now()
	limiterCondition.Active = true
	_, err := repo.collection.Insert(limiterCondition.ID.String(), limiterCondition, nil)
	return limiterCondition.ID, err
}

func (repo *LimiterConditionRepository) Update(id uuid.UUID, limiterCondition *entity.LimiterCondition) error {
	limiterCondition.ID = id
	_, err := repo.collection.Upsert(id.String(), limiterCondition, nil)
	return err
}

func (repo *LimiterConditionRepository) Delete(id uuid.UUID) error {
	_, err := repo.collection.Remove(id.String(), nil)
	return err
}

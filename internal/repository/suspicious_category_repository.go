package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/util"
	"time"
)

type SuspiciousCategoryRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewSuspiciousCategoryRepository(cluster *gocb.Cluster, bucket *gocb.Bucket) *SuspiciousCategoryRepository {
	return &SuspiciousCategoryRepository{
		cluster:    cluster,
		collection: bucket.Collection(suspiciousCategoriesCollection),
	}
}

func (repo *SuspiciousCategoryRepository) FindAllByIds(ids []uuid.UUID) ([]entity.SuspiciousCategory, error) {
	return repo.FindAllByQuery(
		fmt.Sprintf("WHERE meta().id IN [%s]", util.JoinQuotedUUIDs(ids, ",")))
}

func (repo *SuspiciousCategoryRepository) FindAllByQuery(whereClause string) ([]entity.SuspiciousCategory, error) {
	query := fmt.Sprintf(fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), suspiciousCategoriesCollection, whereClause))
	data, err := repo.cluster.Query(query, nil)
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

func (repo *SuspiciousCategoryRepository) CountByQuery(whereClause string) (int, error) {
	query := fmt.Sprintf("SELECT count(meta().id) FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), suspiciousCategoriesCollection, whereClause)
	return countByQuery(repo.cluster, query)
}

func (repo *SuspiciousCategoryRepository) Insert(suspiciousCategory *entity.SuspiciousCategory) (uuid.UUID, error) {
	suspiciousCategory.ID = uuid.New()
	suspiciousCategory.DateTime = time.Now().String()
	_, err := repo.collection.Insert(suspiciousCategory.ID.String(), suspiciousCategory, nil)
	return suspiciousCategory.ID, err
}

func (repo *SuspiciousCategoryRepository) Update(id uuid.UUID, suspiciousCategory *entity.SuspiciousCategory) error {
	suspiciousCategory.ID = id
	_, err := repo.collection.Upsert(id.String(), suspiciousCategory, nil)
	return err
}

func (repo *SuspiciousCategoryRepository) Delete(id uuid.UUID) error {
	_, err := repo.collection.Remove(id.String(), nil)
	return err
}

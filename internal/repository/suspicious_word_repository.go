package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

type SuspiciousWordRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewSuspiciousWordRepository(cluster *gocb.Cluster, bucket *gocb.Bucket) *SuspiciousWordRepository {
	return &SuspiciousWordRepository{
		cluster:    cluster,
		collection: bucket.Collection(suspiciousWordsCollection),
	}
}

func (repo *SuspiciousWordRepository) FindAllByQuery(whereClause string) ([]entity.SuspiciousWord, error) {
	query := fmt.Sprintf(fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), suspiciousWordsCollection, whereClause))
	data, err := repo.cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var result []entity.SuspiciousWord
	for data.Next() {
		var record entity.SuspiciousWord
		if err := data.Row(&record); err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func (repo *SuspiciousWordRepository) CountAllByQuery(whereClause string) (int, error) {
	query := fmt.Sprintf(fmt.Sprintf("SELECT count(meta().id) FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), suspiciousWordsCollection, whereClause))
	return countByQuery(repo.cluster, query)
}

func (repo *SuspiciousWordRepository) Insert(suspiciousWords []entity.SuspiciousWord) error {
	for _, suspiciousWord := range suspiciousWords {
		suspiciousWord.ID = uuid.New()
		suspiciousWord.DateTime = time.Now()
		_, err := repo.collection.Upsert(suspiciousWord.ID.String(), suspiciousWord, nil)
		return err
	}
	return nil
}

func (repo *SuspiciousWordRepository) Delete(id uuid.UUID) error {
	_, err := repo.collection.Remove(id.String(), nil)
	return err
}

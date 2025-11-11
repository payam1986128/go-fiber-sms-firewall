package repository

import (
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/config"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

type SuspiciousWordRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewSuspiciousWordRepository(config *config.CouchbaseConfig) *SuspiciousWordRepository {
	return &SuspiciousWordRepository{
		cluster:    config.Cluster,
		collection: config.Bucket.Collection(suspiciousWordsCollection),
	}
}

func (repo *SuspiciousWordRepository) FindSuspiciousWordsByQuery(query string) ([]entity.SuspiciousWord, error) {
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

func (repo *SuspiciousWordRepository) AddSuspiciousWords(suspiciousWords []entity.SuspiciousWord) error {
	for _, suspiciousWord := range suspiciousWords {
		suspiciousWord.ID = uuid.New()
		suspiciousWord.DateTime = time.Now()
		_, err := repo.collection.Upsert(suspiciousWord.ID.String(), suspiciousWord, nil)
		return err
	}
	return nil
}

func (repo *SuspiciousWordRepository) DeleteSuspiciousWord(id uuid.UUID) error {
	_, err := repo.collection.Remove(id.String(), nil)
	return err
}

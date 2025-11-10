package repository

import (
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

func FindSuspiciousWordsByQuery(query string) ([]entity.SuspiciousWord, error) {
	data, err := db.Cluster.Query(query, nil)
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

func AddSuspiciousWords(suspiciousWords []entity.SuspiciousWord) error {
	for _, suspiciousWord := range suspiciousWords {
		suspiciousWord.ID = uuid.New()
		suspiciousWord.DateTime = time.Now()
		_, err := db.Bucket.Collection(suspiciousWordsCollection).Upsert(suspiciousWord.ID.String(), suspiciousWord, nil)
		return err
	}
	return nil
}

func DeleteSuspiciousWord(id uuid.UUID) error {
	_, err := db.Bucket.Collection(suspiciousWordsCollection).Remove(id.String(), nil)
	return err
}

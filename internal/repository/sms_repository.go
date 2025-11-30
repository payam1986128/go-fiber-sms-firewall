package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

type SmsRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewSmsRepository(cluster *gocb.Cluster, bucket *gocb.Bucket) *SmsRepository {
	return &SmsRepository{
		cluster:    cluster,
		collection: bucket.Collection(smsCollection),
	}
}

func (repo *SmsRepository) Count(sender string, start time.Time, end time.Time) (int, error) {
	return countByQuery(repo.cluster,
		fmt.Sprintf("select count(meta().id) from `%s`.`_default`.`%s` where sender = %s and receivedTime between %s and %s",
			repo.collection.Name(), smsCollection, sender, start.Format(time.RFC3339), end.Format(time.RFC3339)),
	)
}

func (repo *SmsRepository) CountByAppliedFilterId(conditionId uuid.UUID) (int, error) {
	return countByQuery(repo.cluster,
		fmt.Sprintf("select count(meta().id) from `%s`.`_default`.`%s` where appliedFilterId = %s",
			repo.collection.Name(), smsCollection, conditionId),
	)
}

func (repo *SmsRepository) FindAllByQuery(query string) ([]entity.Sms, error) {
	data, err := repo.cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var result []entity.Sms
	for data.Next() {
		var record entity.Sms
		if err := data.Row(&record); err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func (repo *SmsRepository) CountSmsByQuery(whereClause string) (int, error) {
	query := fmt.Sprintf(fmt.Sprintf("SELECT count(meta().id) FROM `%s`.`_default`.`%s` %s", repo.collection.Name(), smsCollection, whereClause))
	return countByQuery(repo.cluster, query)
}

func (repo *SmsRepository) AddSms(sms *entity.Sms) (uuid.UUID, error) {
	sms.ID = uuid.New()
	sms.ReceivedTime = time.Now()
	_, err := repo.collection.Insert(sms.ID.String(), sms, nil)
	return sms.ID, err
}

func (repo *SmsRepository) UpdateSms(sms *entity.Sms) error {
	_, err := repo.collection.Replace(sms.ID.String(), sms, nil)
	return err
}

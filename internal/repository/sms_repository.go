package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

func CountSms(sender string, start time.Time, end time.Time) (int, error) {
	return countByQuery(
		fmt.Sprintf("select count(meta().id) from `%s`.`_default`.`%s` where sender = %s and receivedTime between %s and %s",
			db.Bucket.Name(), smsCollection, sender, start.Format(time.RFC3339), end.Format(time.RFC3339)),
	)
}

func CountSmsByAppliedFilterId(conditionId uuid.UUID) (int, error) {
	return countByQuery(
		fmt.Sprintf("select count(meta().id) from `%s`.`_default`.`%s` where appliedFilterId = %s",
			db.Bucket.Name(), smsCollection, conditionId),
	)
}

func FindSmsByQuery(query string) ([]entity.Sms, error) {
	data, err := db.Cluster.Query(query, nil)
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

func AddSms(sms *entity.Sms) (uuid.UUID, error) {
	sms.ID = uuid.New()
	sms.ReceivedTime = time.Now()
	_, err := db.Bucket.Collection(smsCollection).Insert(sms.ID.String(), sms, nil)
	return sms.ID, err
}

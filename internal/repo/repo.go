package repo

import (
	"fmt"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/models"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
)

func UpsertSMS(s *models.SMS) error {
	if s.ID == "" {
		s.ID = uuid.NewString()
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now().UTC()
	}
	key := "sms::" + s.ID
	_, err := db.Collection.Upsert(key, s, nil)
	return err
}

func GetSMSByID(id string) (*models.SMS, error) {
	key := "sms::" + id
	get, err := db.Collection.Get(key, nil)
	if err != nil {
		return nil, err
	}
	var s models.SMS
	if err := get.Content(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func ListSMSes(limit int) ([]models.SMS, error) {
	if limit <= 0 {
		limit = 50
	}
	q := fmt.Sprintf("SELECT s.* FROM `%s` s WHERE META(s).id LIKE 'sms::%%' LIMIT %d", db.Bucket.Name(), limit)
	rows, err := db.Cluster.Query(q, &gocb.QueryOptions{Adhoc: true})
	if err != nil {
		return nil, err
	}
	var res []models.SMS
	for rows.Next() {
		var r models.SMS
		if err := rows.Row(&r); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func UpsertRule(r *models.Rule) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now().UTC()
	}
	key := "rule::" + r.ID
	_, err := db.Collection.Upsert(key, r, nil)
	return err
}

func GetRuleByID(id string) (*models.Rule, error) {
	key := "rule::" + id
	get, err := db.Collection.Get(key, nil)
	if err != nil {
		return nil, err
	}
	var r models.Rule
	if err := get.Content(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func ListRules(limit int) ([]models.Rule, error) {
	if limit <= 0 {
		limit = 100
	}
	q := fmt.Sprintf("SELECT r.* FROM `%s` r WHERE META(r).id LIKE 'rule::%%' LIMIT %d", db.Bucket.Name(), limit)
	rows, err := db.Cluster.Query(q, &gocb.QueryOptions{Adhoc: true})
	if err != nil {
		return nil, err
	}
	var res []models.Rule
	for rows.Next() {
		var r models.Rule
		if err := rows.Row(&r); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

// Basic search for matching rules in memory. For large scale, push evaluation into queries
func GetAllRules() ([]models.Rule, error) {
	return ListRules(1000)
}

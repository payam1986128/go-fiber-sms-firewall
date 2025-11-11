package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	LimiterCondition struct {
		ID          uuid.UUID   `bson:"_id"`
		Name        string      `bson:"name"`
		Priority    int         `bson:"priority"`
		Active      bool        `bson:"active"`
		CreatedTime time.Time   `bson:"createdTime"`
		TimeLimits  []TimeLimit `bson:"timeLimits"`
		Filters     Filters     `bson:"filters"`
		Action      Action      `bson:"action"`
	}

	Filters struct {
		Keyword   *KeywordsFilter `bson:"keyword"`
		Sender    *SendersFilter  `bson:"sender"`
		Receivers []string        `bson:"receivers"`
	}

	KeywordsFilter struct {
		Keywords         []string    `bson:"keywords"`
		Categories       []uuid.UUID `bson:"categories"`
		CategoryKeywords []string    `bson:"categoryKeywords"`
		Regexes          []string    `bson:"regexes"`
	}

	SendersFilter struct {
		Senders []string    `bson:"senders"`
		Rate    *RateFilter `bson:"rate"`
	}

	RateFilter struct {
		IntervalMinutes int `bson:"intervalMinutes"`
		Threshold       int `bson:"threshold"`
	}

	TimeLimit struct {
		From *time.Time `bson:"from"`
		To   time.Time  `bson:"to"`
	}
)

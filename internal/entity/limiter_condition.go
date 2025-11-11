package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	LimiterCondition struct {
		ID          uuid.UUID `bson:"_id"`
		Name        string
		Priority    int
		Active      bool
		CreatedTime time.Time
		TimeLimits  []TimeLimit
		Filters     Filters
		Action      Action
	}

	Filters struct {
		Keyword   *KeywordsFilter
		Sender    *SendersFilter
		Receivers []string
	}

	KeywordsFilter struct {
		Keywords         []string
		Categories       []uuid.UUID
		CategoryKeywords []string
		Regexes          []string
	}

	SendersFilter struct {
		Senders []string
		Rate    *RateFilter
	}

	RateFilter struct {
		IntervalMinutes int
		Threshold       int
	}

	TimeLimit struct {
		From *time.Time
		To   time.Time
	}
)

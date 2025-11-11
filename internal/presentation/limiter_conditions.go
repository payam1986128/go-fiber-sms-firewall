package presentation

import (
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

type (
	BriefLimiterConditionDto struct {
		ID          string         `json:"id"`
		Name        string         `json:"name"`
		CreatedTime string         `json:"createdTime"`
		Priority    int            `json:"priority"`
		TimeLimits  []TimeLimitDto `json:"timeLimits"`
	}

	TimeLimitDto struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	FiltersDto struct {
		Keyword   *KeywordsFilterDto `json:"keyword"`
		Sender    *SendersFilterDto  `json:"sender"`
		Receivers []string           `json:"receivers"`
	}

	KeywordsFilterDto struct {
		Keywords   []string    `json:"keywords"`
		Categories []uuid.UUID `json:"categories"`
		Regexes    []string    `json:"regexes"`
	}

	SendersFilterDto struct {
		Senders []string       `json:"senders"`
		Rate    *RateFilterDto `json:"rate"`
	}

	RateFilterDto struct {
		IntervalMinutes int `json:"intervalMinutes"`
		Threshold       int `json:"threshold"`
	}

	LimiterConditionRequest struct {
		Name       string         `json:"name" validate:"required"`
		Priority   int            `json:"priority"`
		TimeLimits []TimeLimitDto `json:"timeLimits"`
		Filters    FiltersDto     `json:"filters"`
		Action     entity.Action  `json:"action"`
	}

	LimiterConditionResponse struct {
		Name       string         `json:"name"`
		Priority   int            `json:"priority"`
		Active     bool           `json:"active"`
		CaughtSms  int            `json:"caughtSms"`
		TimeLimits []TimeLimitDto `json:"timeLimits"`
		Filters    FiltersDto     `json:"filters"`
		Action     entity.Action  `json:"action"`
	}

	LimiterConditionsFilterRequest struct {
		Pageable
		Sortable
		State  bool   `json:"state"`
		Filter string `json:"filter"`
	}

	LimiterConditionsResponse struct {
		LimiterConditions []BriefLimiterConditionDto `json:"limiterConditions"`
		Count             int                        `json:"count"`
		SearchTime        time.Time                  `json:"searchTime"`
	}

	LimiterConditionStateRequest struct {
		IDs   []uuid.UUID `json:"ids"`
		State bool        `json:"state"`
	}
)

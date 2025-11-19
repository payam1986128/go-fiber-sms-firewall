package presentation

import (
	"errors"
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
		From *string `json:"from"`
		To   *string `json:"to"`
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
		IntervalMinutes *int `json:"intervalMinutes"`
		Threshold       *int `json:"threshold"`
	}

	LimiterConditionRequest struct {
		Name       *string        `json:"name"`
		Priority   *int           `json:"priority"`
		TimeLimits []TimeLimitDto `json:"timeLimits"`
		Filters    *FiltersDto    `json:"filters"`
		Action     *entity.Action `json:"action"`
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

	LimiterConditionsSearchRequest struct {
		Pageable
		Sortable
		State  *bool   `json:"state"`
		Filter *string `json:"filter"`
	}

	LimiterConditionsResponse struct {
		LimiterConditions []BriefLimiterConditionDto `json:"limiterConditions"`
		Count             int                        `json:"count"`
		SearchTime        time.Time                  `json:"searchTime"`
	}

	LimiterConditionStateRequest struct {
		IDs   []string `json:"ids"`
		State *bool    `json:"state"`
	}
)

func (req *RateFilterDto) Validate() error {
	if req.IntervalMinutes == nil {
		return errors.New("interval minutes is required")
	}
	if *req.IntervalMinutes < 1 {
		return errors.New("interval minutes must be greater than 1")
	}
	if *req.IntervalMinutes < 10 {
		return errors.New("interval minutes must be less than 10")
	}
	if req.Threshold == nil {
		return errors.New("threshold is required")
	}
	if *req.Threshold < 5 {
		return errors.New("threshold must be greater than 5")
	}
	return nil
}

func (req *FiltersDto) Validate() error {
	if req.Keyword == nil && req.Sender == nil && req.Receivers == nil {
		return errors.New("at least one field of `filters.keyword`, `filters.sender`, `filters.receivers` should be filled")
	}
	if req.Keyword != nil {
		if err := req.Keyword.Validate(); err != nil {
			return err
		}
	}
	if req.Sender != nil {
		if err := req.Sender.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (req *KeywordsFilterDto) Validate() error {
	if req.Keywords == nil && req.Categories == nil && req.Regexes == nil {
		return errors.New("at least one field of `filters.keyword.keywords`, `filters.keyword.categories`, `filters.keyword.regexes` should be filled")
	}
	return nil
}

func (req *SendersFilterDto) Validate() error {
	if req.Senders == nil && req.Rate == nil {
		return errors.New("at least one field of `filters.sender.senders`, `filters.sender.rate` should be filled")
	}
	if req.Rate != nil {
		if err := req.Rate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (req *LimiterConditionRequest) Validate() error {
	if req.Name == nil {
		return errors.New("name is required")
	}
	if req.Priority == nil || *req.Priority < 1 {
		return errors.New("priority is required and must be greater than 1")
	}
	if req.Action == nil {
		return errors.New("action is required")
	}
	if req.Filters == nil {
		return errors.New("filters is required")
	}
	if err := req.Filters.Validate(); err != nil {
		return err
	}
	if req.TimeLimits != nil {
		for _, timeLimit := range req.TimeLimits {
			if err := timeLimit.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (req *TimeLimitDto) Validate() error {
	if req.To == nil {
		return errors.New("to is required")
	}
	return nil
}

func (req *LimiterConditionStateRequest) Validate() error {
	if req.IDs == nil || len(req.IDs) == 0 {
		return errors.New("ids are required")
	}
	if req.State == nil {
		return errors.New("state is required")
	}
	return nil
}

package presentation

import (
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"time"
)

type (
	BriefSmsDto struct {
		ID              string        `json:"id"`
		Sender          string        `json:"sender"`
		Receiver        string        `json:"receiver"`
		Message         string        `json:"message"`
		AppliedFilterID string        `json:"appliedFilterId"`
		ReceivedTime    time.Time     `json:"receivedTime"`
		EvaluatedTime   time.Time     `json:"evaluatedTime"`
		SendTime        time.Time     `json:"sendTime"`
		Action          entity.Action `json:"action"`
	}

	SmsFilterRequest struct {
		Pageable
		Sortable
		Sender   string        `json:"sender"`
		Receiver string        `json:"receiver"`
		DateFrom time.Time     `json:"dateFrom"`
		DateTo   time.Time     `json:"dateTo"`
		Action   entity.Action `json:"action"`
	}

	SmsRequest struct {
		ID       uuid.UUID     `json:"id"`
		Sender   string        `json:"sender"`
		Receiver string        `json:"receiver"`
		Message  string        `json:"message"`
		Action   entity.Action `json:"action"`
	}

	SmsResponse struct {
		Sms        []BriefSmsDto `json:"sms"`
		Count      int           `json:"count"`
		SearchTime time.Time     `json:"searchTime"`
	}

	SmsStateRequest struct {
		IDs      []uuid.UUID `json:"ids"`
		Accepted bool        `json:"accepted"`
	}

	SmsTime struct {
		ID           uuid.UUID `json:"id"`
		ReceivedTime time.Time `json:"receivedTime"`
	}
)

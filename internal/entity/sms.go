package entity

import (
	"github.com/google/uuid"
	"time"
)

type Sms struct {
	ID              uuid.UUID `bson:"_id"`
	Sender          string
	Receiver        string
	Message         string
	ReceivedTime    time.Time
	EvaluatedTime   time.Time
	SendTime        time.Time
	AppliedFilterId uuid.UUID
	Action          Action
}

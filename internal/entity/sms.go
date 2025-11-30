package entity

import (
	"github.com/google/uuid"
	"time"
)

type Sms struct {
	ID              uuid.UUID `bson:"_id"`
	Sender          string    `bson:"sender"`
	Receiver        string    `bson:"receiver"`
	Message         string    `bson:"message"`
	ReceivedTime    time.Time `bson:"receivedTime"`
	EvaluatedTime   time.Time `bson:"evaluatedTime"`
	SendTime        time.Time `bson:"sendTime"`
	AppliedFilterId uuid.UUID `bson:"appliedFilterId"`
	Action          Action    `bson:"action"`
}

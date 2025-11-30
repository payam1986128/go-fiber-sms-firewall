package entity

import (
	"github.com/google/uuid"
	"time"
)

type SuspiciousWord struct {
	ID       uuid.UUID          `bson:"_id"`
	Category SuspiciousCategory `bson:"category"`
	Word     string             `bson:"word"`
	DateTime time.Time          `bson:"datetime"`
}

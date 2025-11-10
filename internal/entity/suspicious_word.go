package entity

import (
	"github.com/google/uuid"
	"time"
)

type SuspiciousWord struct {
	ID       uuid.UUID `bson:"_id"`
	Category SuspiciousCategory
	Word     string
	DateTime time.Time
}

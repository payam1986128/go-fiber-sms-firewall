package entity

import (
	"github.com/google/uuid"
)

type SuspiciousCategory struct {
	ID       uuid.UUID `bson:"_id"`
	Name     string
	DateTime string
	Words    []string
}

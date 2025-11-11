package entity

import (
	"github.com/google/uuid"
)

type SuspiciousCategory struct {
	ID       uuid.UUID `bson:"_id"`
	Name     string    `bson:"name"`
	DateTime string    `bson:"datetime"`
	Words    []string  `bson:"words"`
}

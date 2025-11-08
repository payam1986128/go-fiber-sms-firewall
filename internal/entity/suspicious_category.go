package entity

import (
	"github.com/google/uuid"
	"time"
)

type SuspiciousCategory struct {
	ID       uuid.UUID
	name     string
	DateTime time.Time
	Words    []string
}

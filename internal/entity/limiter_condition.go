package entity

import (
	"github.com/google/uuid"
	"time"
)

type LimiterCondition struct {
	ID          uuid.UUID `bson:"_id"`
	Name        string
	Priority    int
	Active      bool
	CreatedTime time.Time
	TimeLimits  []TimeLimit
	Filters     Filters
	Action      Action
}

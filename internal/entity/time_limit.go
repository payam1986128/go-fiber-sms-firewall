package entity

import "time"

type TimeLimit struct {
	from time.Time
	to   time.Time
}

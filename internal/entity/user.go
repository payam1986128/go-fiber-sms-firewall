package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `bson:"_id"`
	Username string
	Password string
	Active   bool
}

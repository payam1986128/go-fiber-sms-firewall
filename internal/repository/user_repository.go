package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
)

func FindUserByUsername(username string) (*entity.User, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` WHERE username = %s", db.Bucket.Name(), userCollection, username)
	data, err := db.Cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var user entity.User
	for data.Next() {
		if err := data.Row(&user); err != nil {
			return nil, err
		}
	}
	if err := data.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUser(user *entity.User) (uuid.UUID, error) {
	user.ID = uuid.New()
	user.Active = true
	_, err := db.Bucket.Collection(userCollection).Insert(user.ID.String(), user, nil)
	return user.ID, err
}

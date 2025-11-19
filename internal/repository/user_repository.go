package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/config"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
)

type UserRepository struct {
	cluster    *gocb.Cluster
	collection *gocb.Collection
}

func NewUserRepository(config *config.CouchbaseConfig) *UserRepository {
	return &UserRepository{
		cluster:    config.Cluster,
		collection: config.Bucket.Collection(userCollection),
	}
}

func (repo *UserRepository) FindByUsername(username string) (*entity.User, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`.`_default`.`%s` WHERE username = %s", repo.collection.Name(), userCollection, username)
	data, err := repo.cluster.Query(query, nil)
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

func (repo *UserRepository) Insert(user *entity.User) (uuid.UUID, error) {
	user.ID = uuid.New()
	user.Active = true
	_, err := repo.collection.Insert(user.ID.String(), user, nil)
	return user.ID, err
}

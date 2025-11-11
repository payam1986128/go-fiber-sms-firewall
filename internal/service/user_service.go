package service

import "github.com/payam1986128/go-fiber-sms-firewall/internal/repository"

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

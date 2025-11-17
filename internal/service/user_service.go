package service

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (service *UserService) RegisterUser(request *presentation.RegisterUserRequest, code string) (string, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	user := &entity.User{
		Username: *request.Username,
		Password: string(hashed),
		Active:   true,
	}
	id, err := service.repository.AddUser(user)
	return id.String(), err
}

func (service *UserService) LoginUser(username string, password string) (*presentation.VerificationResponse, error) {
	return nil, nil
}

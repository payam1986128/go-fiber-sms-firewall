package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService struct {
	repository *repository.UserRepository
	secret     string
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
		secret:     os.Getenv("JWT_SECRET"),
	}
}

func (service *UserService) RegisterUser(request *presentation.RegisterUserRequest, code string) (string, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	user := &entity.User{
		Username: *request.Username,
		Password: string(hashed),
		Active:   true,
	}
	id, err := service.repository.Insert(user)
	return id.String(), err
}

func (service *UserService) LoginUser(username string, code string) (*presentation.VerificationResponse, error) {
	user, err := service.repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(code)) != nil {
		return nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte(service.secret))
	return &presentation.VerificationResponse{
		AccessToken: signed,
	}, err
}

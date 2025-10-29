package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user string) (string, error) {
	secret := []byte(getJWTSecret())
	claims := jwt.MapClaims{
		"sub": user,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func getJWTSecret() string {
	s := "changeme"
	if v := os.Getenv("JWT_SECRET"); v != "" {
		s = v
	}
	return s
}

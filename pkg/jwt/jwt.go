package jwt

import (
	"errors"
	"time"

	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrExpired       = errors.New("token expired")
	ErrInvalidClaims = errors.New("invalid claims")
)

func NewToken(user models.User, secret string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	return token.SignedString([]byte(secret))
}

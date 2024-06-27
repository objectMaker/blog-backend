package jwt

import (
	"fmt"
	"os"
	"time"

	v5 "github.com/golang-jwt/jwt/v5"
)

func New(username string) (string, error) {
	tS := os.Getenv("TOKEN_SECRET")
	if tS == "" {
		return "", fmt.Errorf("TOKEN_SECRET is empty")
	}
	token := v5.NewWithClaims(v5.SigningMethodHS256, v5.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).UnixNano(),
	})
	tokenString, err := token.SignedString(tS)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return tokenString, nil
}

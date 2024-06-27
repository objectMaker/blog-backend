package jwt

import (
	"errors"
	"fmt"
	"log"
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
	tokenString, err := token.SignedString([]byte(tS))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return tokenString, nil
}

type Payload struct {
	Username string  `json:"username"`
	Exp      float64 `json:"exp"`
}

func ParseToken(tokenString string) (Payload, error) {
	token, err := v5.Parse(tokenString, func(token *v5.Token) (interface{}, error) {
		if _, ok := token.Method.(*v5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(v5.MapClaims); ok {
		return Payload{
			Username: claims["username"].(string),
			Exp:      claims["exp"].(float64),
		}, nil
	}
	return Payload{}, errors.New("parse token failed")
}

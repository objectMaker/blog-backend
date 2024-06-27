package jwt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	v5 "github.com/golang-jwt/jwt/v5"
	"github.com/objectMaker/blog-backend/tools"
)

func New(username string) (string, error) {
	tS := os.Getenv("TOKEN_SECRET")
	if tS == "" {
		return "", fmt.Errorf("TOKEN_SECRET is empty")
	}
	tokenExp, err := tools.GetEnvInt("TOKEN_EXPIRATION")
	if err != nil {
		fmt.Sprintln("ERROR: %w", err)
		return "", err

	}

	token := v5.NewWithClaims(v5.SigningMethodHS256, v5.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(tokenExp)).Unix(),
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

func ValidateToken(token string) (bool, error) {
	payload, err := ParseToken(token)
	if err != nil {
		return false, err
	}
	//
	if payload.Exp < float64(time.Now().UnixNano()) {
		return false, fmt.Errorf("token expired")
	}
	return true, nil
}

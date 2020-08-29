package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewUserClaims(id int, name string) UserClaims {
	uc := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3600 * time.Second).Unix(),
			Issuer:    "jdxj",
		},
		ID:   id,
		Name: name,
	}
	return uc
}

type UserClaims struct {
	jwt.StandardClaims

	ID   int    `json:"id"`
	Name string `json:"name"`
}

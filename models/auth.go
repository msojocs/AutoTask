package model

import "github.com/golang-jwt/jwt/v4"

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

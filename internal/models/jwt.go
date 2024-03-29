package models

import "github.com/golang-jwt/jwt/v5"

type JWTContent struct {
	Mail string `json:"mail"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

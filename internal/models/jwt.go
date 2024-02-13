package models

import "github.com/golang-jwt/jwt/v5"

type JWTContent struct {
	Id   string `json:"id"`
	Mail string `json:"mail"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

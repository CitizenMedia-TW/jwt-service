package models

import "github.com/golang-jwt/jwt/v5"

type JWTContent struct {
	Id   string
	Mail string
	Name string
	jwt.RegisteredClaims
}

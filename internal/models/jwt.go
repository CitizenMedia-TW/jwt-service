package models

import "github.com/golang-jwt/jwt/v5"

type JWTContent struct {
	UserMail string
	UserName string
	jwt.RegisteredClaims
}

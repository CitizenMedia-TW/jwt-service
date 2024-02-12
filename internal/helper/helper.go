package helper

import (
	"auth-service/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWTSignContent(content models.JWTContent, key string) (string, error) {
	// Set up the claims
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)), // 2 day expiration
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	content.RegisteredClaims = claims

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, content)
	// Sign the token with the key
	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JWTParseToken(tokenString string, key string) (*models.JWTContent, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTContent{}, keyFunc)
	if err != nil {
		return nil, err
	}

	// Validate the token and return the claims
	if claims, ok := token.Claims.(*models.JWTContent); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

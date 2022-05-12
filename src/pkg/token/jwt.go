package token

import (
	"github.com/golang-jwt/jwt"
)

type JWTToken struct {
}

func NewJWTToken() Token {
	return &JWTToken{}
}

// Generate generates a access token
func (j *JWTToken) Generate(payload map[string]interface{}, jwtKey string) (string, error) {
	claims := jwt.MapClaims{}
	claims = payload

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Validate is method to check if a token is valid or not
func (j *JWTToken) Validate(token string, jwtKey string) (map[string]interface{}, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, nil
	}

	return claims, nil
}

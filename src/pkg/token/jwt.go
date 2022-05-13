package token

import (
	"github.com/golang-jwt/jwt"
)

type JWTToken struct {
	jwtKey string
}

func NewJWTToken(key string) Token {
	return &JWTToken{
		jwtKey: key,
	}
}

// Generate generates a access token
func (j *JWTToken) Generate(payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}
	claims = payload

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(j.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Validate is method to check if a token is valid or not
func (j *JWTToken) Validate(token string) (map[string]interface{}, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(j.jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, nil
	}

	return claims, nil
}

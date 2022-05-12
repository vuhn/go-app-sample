package service

// AccessTokenService service interface
type AccessTokenService interface {
	GenerateToken(claims map[string]interface{}, secretKey string) (string, error)
	ValidateToken(token string, key string) (bool, error)
}

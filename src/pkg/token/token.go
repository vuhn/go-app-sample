package token

// Token is a interface to define methods for package token
type Token interface {
	Generate(claims map[string]interface{}, secretKey string) (string, error)
	Validate(token string, key string) (map[string]interface{}, error)
}

package token

// Token is a interface to define methods for package token
type Token interface {
	Generate(claims map[string]interface{}) (string, error)
	Validate(token string) (map[string]interface{}, error)
}

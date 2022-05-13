package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var idLength = 36

type JWTTokenTestSuite struct {
	suite.Suite
	jwtToken Token
}

func (j *JWTTokenTestSuite) SetupTest() {
	jtwKey := "123456"
	jwtToken := NewJWTToken(jtwKey)
	j.jwtToken = jwtToken
}
func (j *JWTTokenTestSuite) TestGenerate_ShouldReturnSuccess() {
	claims := map[string]interface{}{
		"id": "123456",
	}
	token, err := j.jwtToken.Generate(claims)
	j.Nil(err)
	j.NotEmpty(token)
}

func (j *JWTTokenTestSuite) TestValidate_ShouldReturnSuccess() {
	id := "123456"
	claims := map[string]interface{}{
		"id": id,
	}
	token, err := j.jwtToken.Generate(claims)
	j.Nil(err)
	j.NotEmpty(token)

	result, err := j.jwtToken.Validate(token)
	j.Nil(err)
	j.NotNil(result)
	j.Equal(id, result["id"])
}

func (j *JWTTokenTestSuite) TestValidate_ShouldReturnError() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZ1aG41QGdtYWlsLmNvbSIsInVzZXJfaWQiOiI0ODI5ODU0OS1hYzRjLTQzMWYtOTc4OC03ZjE3OWQ2Y2QzZjEifQ.s7TI9qzvPZubWI_-qlwPevXcKVv3NAZV1A2kC8GXc0s"

	result, err := j.jwtToken.Validate(token)
	j.NotNil(err)
	j.Nil(result)
}

func (j *JWTTokenTestSuite) TestValidate_ShouldReturnErrorWhenExpired() {
	id := "123456"
	claims := map[string]interface{}{
		"id":  id,
		"exp": time.Now(),
	}
	token, err := j.jwtToken.Generate(claims)
	j.Nil(err)
	j.NotEmpty(token)

	result, err := j.jwtToken.Validate(token)
	j.Error(err)
	j.Nil(result)
}

func TestGenerateNewID(t *testing.T) {
	suite.Run(t, new(JWTTokenTestSuite))
}

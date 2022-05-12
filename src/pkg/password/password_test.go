package password

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var encryptedPasswordLength = 60

type PasswordTestSuite struct {
	suite.Suite
	password Password
}

func (p *PasswordTestSuite) SetupTest() {
	password := NewBcryptPassword()
	p.password = password
}
func (p *PasswordTestSuite) TestGenerateNewPassword_ShouldReturnEncryptedPassword() {
	plainPassword := "123456"
	encryptedPassword, err := p.password.GenerateFromPassword(plainPassword)
	p.NoError(err)
	p.Len(encryptedPassword, encryptedPasswordLength)
}

func (p *PasswordTestSuite) TestCompareHashAndPassword_ShouldReturnSame() {
	plainPassword := "123456"
	encryptedPassword, err := p.password.GenerateFromPassword(plainPassword)
	p.NoError(err)
	p.Len(encryptedPassword, encryptedPasswordLength)

	result := p.password.CompareHashAndPassword(encryptedPassword, plainPassword)
	p.Equal(true, result)
}

func TestGenerateNewID(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}

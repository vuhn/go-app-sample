package password

import "golang.org/x/crypto/bcrypt"

type Password interface {
	GenerateFromPassword(plainPassword string) (string, error)
	CompareHashAndPassword(plainPassword string, hashedPassword string) bool
}

type bcryptPassword struct{}

func NewBcryptPassword() Password {
	return &bcryptPassword{}
}

func (p *bcryptPassword) GenerateFromPassword(plainPassword string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

func (p *bcryptPassword) CompareHashAndPassword(hashedPassword string, plainPassword string) bool {
	plainPasswordBytes := []byte(plainPassword)
	hashedPasswordBytes := []byte(hashedPassword)

	if err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, plainPasswordBytes); err != nil {
		return false
	}

	return true
}

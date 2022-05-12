package idgenerator

import "github.com/google/uuid"

// IDGenerator is id generator interface
type IDGenerator interface {
	GenerateNewID() string
}

// NewIDGenerator returns implementation of IDGenerator interfface
func NewIDGenerator() IDGenerator {
	return &uuidGenerator{}
}

type uuidGenerator struct {
}

func (d *uuidGenerator) GenerateNewID() string {
	return uuid.NewString()
}

package idgenerator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var idLength = 36

type UUIDGeneratorTestSuite struct {
	suite.Suite
	idGenerator IDGenerator
}

func (s *UUIDGeneratorTestSuite) SetupTest() {
	idGenerator := NewIDGenerator()
	s.idGenerator = idGenerator
}
func (s *UUIDGeneratorTestSuite) TestGenerateNewID_ShouldCreateNewUUID() {
	id := s.idGenerator.GenerateNewID()
	s.Len(id, idLength)
}

func (s *UUIDGeneratorTestSuite) TestGenerateNewID_ShouldCreate2DiffirentUUIDs() {
	id1 := s.idGenerator.GenerateNewID()
	s.Len(id1, idLength)

	id2 := s.idGenerator.GenerateNewID()
	s.Len(id2, idLength)

	s.NotEqual(id1, id2)
}
func TestGenerateNewID(t *testing.T) {
	suite.Run(t, new(UUIDGeneratorTestSuite))
}

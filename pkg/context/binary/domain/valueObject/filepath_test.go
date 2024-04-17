package valueObject_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/domain/errors"
	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type FilepathValueObjectTestSuite struct {
	suite.Suite
}

func (suite *FilepathValueObjectTestSuite) SetupTest() {}

func (suite *FilepathValueObjectTestSuite) TestFilepath() {
	expected := errors.InvalidValue{Message: "Filepath Invalid"}

	suite.PanicsWithValue(expected, func() { valueObjectMother.InvalidFilepath() })
	suite.PanicsWithValue(expected, func() { valueObjectMother.EmptyFilepath() })
}

func TestUnitFilepathValueObjectSuite(t *testing.T) {
	suite.Run(t, new(FilepathValueObjectTestSuite))
}

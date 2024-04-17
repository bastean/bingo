package valueObject_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/domain/errors"
	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type FilenameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *FilenameValueObjectTestSuite) SetupTest() {}

func (suite *FilenameValueObjectTestSuite) TestFilename() {
	expected := errors.InvalidValue{Message: "Filename must be between " + "1" + " to " + "50" + " characters"}

	suite.PanicsWithValue(expected, func() { valueObjectMother.WithInvalidFilenameLength() })
	suite.PanicsWithValue(expected, func() { valueObjectMother.WithInvalidFilenameAlphanumeric() })
}

func (suite *FilenameValueObjectTestSuite) TestFilenameDefault() {
	expected := "bingo"

	actual := valueObjectMother.EmptyFilename().Value

	suite.EqualValues(expected, actual)
}

func TestUnitFilenameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(FilenameValueObjectTestSuite))
}

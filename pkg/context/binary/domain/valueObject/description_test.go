package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type DescriptionValueObjectTestSuite struct {
	suite.Suite
}

func (suite *DescriptionValueObjectTestSuite) SetupTest() {}

func (suite *DescriptionValueObjectTestSuite) TestDescription() {
	expected := errors.InvalidValue{Message: "Description must be between " + "1" + " to " + "100" + " characters"}

	suite.PanicsWithValue(expected, func() { valueObjectMother.WithInvalidDescriptionLength() })
	suite.PanicsWithValue(expected, func() { valueObjectMother.WithInvalidDescriptionAlphanumeric() })
}

func (suite *DescriptionValueObjectTestSuite) TestDescriptionDefault() {
	expected := "Example service for generating custom cross-compile executable binaries"

	actual := valueObjectMother.EmptyDescription().Value

	suite.EqualValues(expected, actual)
}

func TestUnitDescriptionValueObjectSuite(t *testing.T) {
	suite.Run(t, new(DescriptionValueObjectTestSuite))
}

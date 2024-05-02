package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type OSValueObjectTestSuite struct {
	suite.Suite
}

func (suite *OSValueObjectTestSuite) SetupTest() {}

func (suite *OSValueObjectTestSuite) TestOS() {
	expected := errors.InvalidValue{Message: "Operating System should be: linux, windows, darwin"}

	suite.PanicsWithValue(expected, func() { valueObjectMother.InvalidOS() })
}

func (suite *OSValueObjectTestSuite) TestOSDefault() {
	expected := "linux"

	actual := valueObjectMother.EmptyOS().Value

	suite.EqualValues(expected, actual)
}

func TestUnitOSValueObjectSuite(t *testing.T) {
	suite.Run(t, new(OSValueObjectTestSuite))
}

package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type ArchValueObjectTestSuite struct {
	suite.Suite
}

func (suite *ArchValueObjectTestSuite) SetupTest() {}

func (suite *ArchValueObjectTestSuite) TestArch() {
	expected := errors.InvalidValue{Message: "Architecture should be: amd64, 386"}

	suite.PanicsWithValue(expected, func() { valueObjectMother.InvalidArch() })
}

func (suite *ArchValueObjectTestSuite) TestArchDefault() {
	expected := "amd64"

	actual := valueObjectMother.EmptyArch().Value

	suite.EqualValues(expected, actual)
}

func TestUnitArchValueObjectSuite(t *testing.T) {
	suite.Run(t, new(ArchValueObjectTestSuite))
}

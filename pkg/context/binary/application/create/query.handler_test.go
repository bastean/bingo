package create_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/application/create"
	"github.com/stretchr/testify/suite"
)

type BinaryCreateTestSuite struct {
	suite.Suite
	sut    *create.QueryHandler
	create *create.Create
}

func (suite *BinaryCreateTestSuite) SetupTest() {
	suite.create = create.NewCreate()
	suite.sut = create.NewQueryHandler(suite.create)
}

func (suite *BinaryCreateTestSuite) TestCreate() {
	name := "success"

	description := "success"

	query := create.NewQuery(name, description)

	expected := create.NewResponse("build/success")

	actual := suite.sut.Handle(query)

	suite.EqualValues(expected, actual)
}

func TestUnitBinaryCreateSuite(t *testing.T) {
	suite.Run(t, new(BinaryCreateTestSuite))
}

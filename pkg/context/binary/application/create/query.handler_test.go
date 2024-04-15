package create_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/application/create"
	compilerMock "github.com/bastean/bingo/pkg/context/binary/infrastructure/compiler/mock"
	"github.com/stretchr/testify/suite"
)

type BinaryCreateTestSuite struct {
	suite.Suite
	sut     *create.QueryHandler
	create  *create.Create
	builder *compilerMock.BuilderMock
}

func (suite *BinaryCreateTestSuite) SetupTest() {
	suite.builder = compilerMock.NewBuilderMock()
	suite.create = create.NewCreate(suite.builder)
	suite.sut = create.NewQueryHandler(suite.create)
}

func (suite *BinaryCreateTestSuite) TestCreate() {
	name := "success"

	description := "success"

	query := create.NewQuery(name, description)

	// TODO!: suite.builder.On("build")

	expected := create.NewResponse("build/success")

	actual := suite.sut.Handle(query)

	suite.EqualValues(expected, actual)
}

func TestUnitBinaryCreateSuite(t *testing.T) {
	suite.Run(t, new(BinaryCreateTestSuite))
}

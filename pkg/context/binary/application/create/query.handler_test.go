package create_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/application/create"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
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

	filepath := valueObject.NewFilepath("/build/success")

	query := create.NewQuery(name, description)

	// TODO?: root := aggregate.NewRoot(name, description)

	// TODO?: service.Switcher(root.Command, query.Command)

	// TODO?: suite.builder.On("Build", root).Return(filepath)

	suite.builder.On("Build").Return(filepath)

	expected := filepath.Value

	response := suite.sut.Handle(query)

	actual := response.FilePath

	suite.builder.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitBinaryCreateSuite(t *testing.T) {
	suite.Run(t, new(BinaryCreateTestSuite))
}

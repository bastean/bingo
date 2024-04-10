package builder_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/fakeit/application/builder"
	"github.com/stretchr/testify/suite"
)

type FakeitBuilderTestSuite struct {
	suite.Suite
	sut     *builder.QueryHandler
	builder *builder.Builder
}

func (suite *FakeitBuilderTestSuite) SetupTest() {
	suite.builder = builder.NewBuilder()
	suite.sut = builder.NewQueryHandler(suite.builder)
}

func (suite *FakeitBuilderTestSuite) TestBuilder() {
	// TODO: fakeit builder query and response handler test

	query := builder.NewQuery()

	expected := builder.NewResponse("success")

	actual := suite.sut.Handle(query)

	suite.EqualValues(expected, actual)
}

func TestUnitFakeitBuilderSuite(t *testing.T) {
	suite.Run(t, new(FakeitBuilderTestSuite))
}

package compiler_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/infrastructure/compiler"
	"github.com/stretchr/testify/suite"
)

type BinaryCLIBuilderTestSuite struct {
	suite.Suite
	sut model.Builder
}

func (suite *BinaryCLIBuilderTestSuite) SetupTest() {
	suite.sut = compiler.NewBinaryCLIBuilder()
}

func (suite *BinaryCLIBuilderTestSuite) TestBuild() {}

func TestIntegrationBinaryCLIBuilderSuite(t *testing.T) {
	suite.Run(t, new(BinaryCLIBuilderTestSuite))
}

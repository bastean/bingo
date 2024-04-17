package compiler_test

import (
	"testing"

	"github.com/bastean/bingo/pkg/context/binary/domain/builder"
	"github.com/bastean/bingo/pkg/context/binary/infrastructure/compiler"
	"github.com/stretchr/testify/suite"
)

type BinaryBuilderTestSuite struct {
	suite.Suite
	sut builder.Builder
}

func (suite *BinaryBuilderTestSuite) SetupTest() {
	suite.sut = compiler.NewBinaryBuilder()
}

func (suite *BinaryBuilderTestSuite) TestBuild() {}

func TestIntegrationBinaryBuilderSuite(t *testing.T) {
	suite.Run(t, new(BinaryBuilderTestSuite))
}

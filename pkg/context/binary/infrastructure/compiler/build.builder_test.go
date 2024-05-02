package compiler_test

import (
	"os"
	"os/exec"
	"testing"

	aggregateMother "github.com/bastean/bingo/pkg/context/binary/domain/aggregate/mother"
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

func (suite *BinaryBuilderTestSuite) TestBuild() {
	binary := aggregateMother.Empty()

	expected := `Example service for generating cross-platform custom executable binary

Usage:
  bingo [flags]

Flags:
  -h, --help   help for bingo
`

	filepath := suite.sut.Build(binary).Value

	actual, _ := exec.Command("./" + filepath).Output()

	suite.EqualValues(expected, actual)
}

func (suite *BinaryBuilderTestSuite) TearDownTest() {
	os.RemoveAll("temp")
}

func TestIntegrationBinaryBuilderSuite(t *testing.T) {
	suite.Run(t, new(BinaryBuilderTestSuite))
}

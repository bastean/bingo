package compilerMock

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
	"github.com/stretchr/testify/mock"
)

type BuilderMock struct {
	mock.Mock
}

func (mock *BuilderMock) Build(root *aggregate.Root) *valueObject.FilePath {
	args := mock.Called(root)
	return args.Get(0).(*valueObject.FilePath)
}

func NewBuilderMock() *BuilderMock {
	return new(BuilderMock)
}

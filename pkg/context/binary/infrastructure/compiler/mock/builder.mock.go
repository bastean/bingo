package compilerMock

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
	"github.com/stretchr/testify/mock"
)

type BuilderMock struct {
	mock.Mock
}

func (mock *BuilderMock) Build(binary *aggregate.Binary) *valueObject.Filepath {
	// TODO?: args := mock.Called(binary)

	args := mock.Called()

	return args.Get(0).(*valueObject.Filepath)
}

func NewBuilderMock() *BuilderMock {
	return new(BuilderMock)
}

package compilerMock

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
	"github.com/stretchr/testify/mock"
)

type BuilderMock struct {
	mock.Mock
}

func (mock *BuilderMock) Build(record *model.Record) *valueObject.FilePath {
	args := mock.Called(record)
	return args.Get(0).(*valueObject.FilePath)
}

func NewBuilderMock() *BuilderMock {
	return new(BuilderMock)
}

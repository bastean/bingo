package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Create struct {
	model.Builder
}

func (create *Create) Run(binary *aggregate.Binary) *valueObject.FilePath {
	filePath := create.Builder.Build(binary.Record)

	return filePath
}

func NewCreate(builder model.Builder) *Create {
	return &Create{
		Builder: builder,
	}
}

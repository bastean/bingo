package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/builder"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Create struct {
	builder.Builder
}

func (create *Create) Run(binary *aggregate.Binary) *valueObject.Filepath {
	filePath := create.Builder.Build(binary)

	return filePath
}

func NewCreate(builder builder.Builder) *Create {
	return &Create{
		Builder: builder,
	}
}

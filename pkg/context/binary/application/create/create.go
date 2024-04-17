package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/builder"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Create struct {
	builder.Builder
}

func (create *Create) Run(root *aggregate.Root) *valueObject.FilePath {
	filePath := create.Builder.Build(root)

	return filePath
}

func NewCreate(builder builder.Builder) *Create {
	return &Create{
		Builder: builder,
	}
}

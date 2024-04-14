package create

import (
	"fmt"

	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Create struct{}

func (create *Create) Run(binary *aggregate.Binary) *valueObject.FilePath {
	// TODO: binary create use case

	path := fmt.Sprintf("build/%s", binary.Name)

	return valueObject.NewFilePath(path)
}

func NewCreate() *Create {
	return new(Create)
}

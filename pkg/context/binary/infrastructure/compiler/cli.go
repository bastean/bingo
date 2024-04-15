package compiler

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type CLI struct{}

func (cli *CLI) Build(record *model.Record) *valueObject.FilePath {
	return valueObject.NewFilePath("")
}

func NewBinaryCLIBuilder() model.Builder {
	return new(CLI)
}

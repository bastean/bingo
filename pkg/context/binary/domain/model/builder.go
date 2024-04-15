package model

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Builder interface {
	Build(record *Record) *valueObject.FilePath
}

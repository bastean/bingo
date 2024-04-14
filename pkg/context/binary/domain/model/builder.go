package model

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Builder interface {
	AddRecord(record *Record)
	Build() *valueObject.FilePath
}

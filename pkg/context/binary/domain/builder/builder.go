package builder

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Builder interface {
	Build(binary *aggregate.Binary) *valueObject.Filepath
}

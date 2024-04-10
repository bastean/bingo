package builder

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/aggregate"
)

type Builder struct{}

func (builder *Builder) Run(fakeit *aggregate.Fakeit) {
	// TODO: fakeit builder use case
}

func NewBuilder() *Builder {
	return new(Builder)
}

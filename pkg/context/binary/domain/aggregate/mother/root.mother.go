package aggregateMother

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
)

func Random() *aggregate.Root {
	return aggregate.NewRoot(valueObjectMother.RandomFilename().Value, valueObjectMother.RandomDescription().Value)
}

func Invalid() *aggregate.Root {
	return aggregate.NewRoot(valueObjectMother.WithInvalidFilenameLength().Value, valueObjectMother.WithInvalidDescriptionLength().Value)
}

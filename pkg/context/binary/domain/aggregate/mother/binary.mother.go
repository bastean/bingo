package aggregateMother

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
)

func Random() *aggregate.Binary {
	return aggregate.NewBinary(valueObjectMother.RandomOS().Value, valueObjectMother.RandomArch().Value, valueObjectMother.RandomFilename().Value, valueObjectMother.RandomDescription().Value)
}

func Invalid() *aggregate.Binary {
	return aggregate.NewBinary(valueObjectMother.InvalidOS().Value, valueObjectMother.InvalidArch().Value, valueObjectMother.WithInvalidFilenameLength().Value, valueObjectMother.WithInvalidDescriptionLength().Value)
}

func Empty() *aggregate.Binary {
	return aggregate.NewBinary("", "", "", "")
}

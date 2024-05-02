package queryMother

import (
	"github.com/bastean/bingo/pkg/context/binary/application/create"
	valueObjectMother "github.com/bastean/bingo/pkg/context/binary/domain/valueObject/mother"
)

func Random() *create.Query {
	return create.NewQuery(valueObjectMother.RandomOS().Value, valueObjectMother.RandomArch().Value, valueObjectMother.RandomFilename().Value, valueObjectMother.RandomDescription().Value)
}

func Invalid() *create.Query {
	return create.NewQuery(valueObjectMother.InvalidOS().Value, valueObjectMother.InvalidArch().Value, valueObjectMother.WithInvalidDescriptionLength().Value, valueObjectMother.WithInvalidDescriptionLength().Value)
}

package valueObjectMother

import (
	"strings"

	"github.com/bastean/bingo/pkg/context/binary/domain/service/mother"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

func RandomDescription() *valueObject.Description {
	return valueObject.NewDescription(mother.Create.Regex(`[\w\s.:-]{1,100}`))
}

func WithInvalidDescriptionLength() *valueObject.Description {
	return valueObject.NewDescription(strings.Repeat("x", 102))
}

func WithInvalidDescriptionAlphanumeric() *valueObject.Description {
	return valueObject.NewDescription("<></>")
}

func EmptyDescription() *valueObject.Description {
	return valueObject.NewDescription("")
}

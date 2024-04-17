package valueObjectMother

import (
	"strings"

	"github.com/bastean/bingo/pkg/context/binary/domain/service/mother"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

func RandomFilename() *valueObject.Filename {
	return valueObject.NewFilename(mother.Create.Regex(`[\w-]{1,50}`))
}

func WithInvalidFilenameLength() *valueObject.Filename {
	return valueObject.NewFilename(strings.Repeat("x", 51))
}

func WithInvalidFilenameAlphanumeric() *valueObject.Filename {
	return valueObject.NewFilename("<></>")
}

func EmptyFilename() *valueObject.Filename {
	return valueObject.NewFilename("")
}

package valueObject

import (
	"regexp"
	"strings"

	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
)

const FilenameMinCharactersLength = "1"
const FilenameMaxCharactersLength = "50"

var InvalidFilenameValue = errors.InvalidValue{Message: "Filename must be between " + FilenameMinCharactersLength + " to " + FilenameMaxCharactersLength + " characters"}

type Filename struct {
	Value string
}

func ensureIsValidFilename(filename string) {
	validate := regexp.MustCompile(`^[\w-]{1,50}$`)

	if !validate.MatchString(filename) {
		panic(InvalidFilenameValue)
	}
}

func NewFilename(filename string) *Filename {
	if filename == "" {
		filename = "bingo"
	}

	filename = strings.TrimSpace(filename)

	ensureIsValidFilename(filename)

	filenameVO := &Filename{filename}

	return filenameVO
}

package valueObject

import (
	"regexp"
	"strings"

	"github.com/bastean/bingo/pkg/context/binary/domain/errors"
)

var InvalidFilepathValue = errors.InvalidValue{Message: "Filepath Invalid"}

type Filepath struct {
	Value string
}

func ensureIsValidFilepath(filepath string) {
	validate := regexp.MustCompile(`^\/{0,1}([\w-.()]{1,}\/{0,1})+$`)

	if !validate.MatchString(filepath) {
		panic(InvalidFilepathValue)
	}
}

func NewFilepath(filepath string) *Filepath {
	filepath = strings.TrimSpace(filepath)

	ensureIsValidFilepath(filepath)

	filepathVO := &Filepath{filepath}

	return filepathVO
}

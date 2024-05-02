package valueObject

import (
	"slices"
	"strings"

	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
)

var AllowedOS = []string{"linux", "windows", "darwin"}

var InvalidOSValue = errors.InvalidValue{Message: "Operating System should be: " + strings.Join(AllowedOS, ", ")}

type OS struct {
	Value string
}

func ensureIsValidOS(os string) {
	if !slices.Contains(AllowedOS, os) {
		panic(InvalidOSValue)
	}
}

func NewOS(os string) *OS {
	if os == "" {
		os = "linux"
	}

	os = strings.TrimSpace(os)

	ensureIsValidOS(os)

	OSVO := &OS{os}

	return OSVO
}

package valueObject

import (
	"slices"
	"strings"

	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
)

var AllowedArch = []string{"amd64", "386"}

var InvalidArchValue = errors.InvalidValue{Message: "Architecture should be: " + strings.Join(AllowedArch, ", ")}

type Arch struct {
	Value string
}

func ensureIsValidArch(arch string) {
	if !slices.Contains(AllowedArch, arch) {
		panic(InvalidArchValue)
	}
}

func NewArch(arch string) *Arch {
	if arch == "" {
		arch = "amd64"
	}

	arch = strings.TrimSpace(arch)

	ensureIsValidArch(arch)

	ArchVO := &Arch{arch}

	return ArchVO
}

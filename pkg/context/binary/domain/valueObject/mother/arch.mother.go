package valueObjectMother

import (
	"math/rand/v2"

	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

func RandomArch() *valueObject.Arch {
	randomPick := rand.IntN(len(valueObject.AllowedArch))

	return valueObject.NewArch(valueObject.AllowedArch[randomPick])
}

func InvalidArch() *valueObject.Arch {
	return valueObject.NewArch("x")
}

func EmptyArch() *valueObject.Arch {
	return valueObject.NewArch("")
}

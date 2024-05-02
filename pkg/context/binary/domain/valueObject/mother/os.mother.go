package valueObjectMother

import (
	"math/rand/v2"

	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

func RandomOS() *valueObject.OS {
	randomPick := rand.IntN(len(valueObject.AllowedOS))

	return valueObject.NewOS(valueObject.AllowedOS[randomPick])
}

func InvalidOS() *valueObject.OS {
	return valueObject.NewOS("x")
}

func EmptyOS() *valueObject.OS {
	return valueObject.NewOS("")
}

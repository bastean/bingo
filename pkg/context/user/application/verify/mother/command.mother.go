package commandMother

import (
	"github.com/bastean/bingo/pkg/context/user/application/verify"
	valueObjectMother "github.com/bastean/bingo/pkg/context/user/domain/valueObject/mother"
)

func Random() *verify.Command {
	return verify.NewCommand(valueObjectMother.RandomId().Value)
}

func Invalid() *verify.Command {
	return verify.NewCommand(valueObjectMother.InvalidId().Value)
}

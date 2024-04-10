package aggregate

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/command"
)

type Fakeit struct {
	*command.Person
	*command.Address
	*command.Payment
	*command.Internet
	*command.Language
	*command.Animal
}

func NewFakeit() *Fakeit {
	return new(Fakeit)
}

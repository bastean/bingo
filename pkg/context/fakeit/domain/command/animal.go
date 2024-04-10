package command

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Animal struct {
	*model.Command
}

func NewAnimal() *Animal {
	return &Animal{
		&model.Command{
			Name:        "animal",
			Description: "Generate random animal information",
			Arguments: []*model.Argument{
				{Name: "petName", Description: "Generate random pet name"},
				{Name: "cat", Description: "Generate random cat"},
				{Name: "dog", Description: "Generate random dog"},
				{Name: "bird", Description: "Generate random bird"},
			},
		},
	}
}

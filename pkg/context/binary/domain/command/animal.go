package command

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewAnimal() *model.Command {
	return &model.Command{
		Data: &model.Data{
			Name:        "animal",
			Description: "Generate random animal information",
		},
		Flags: []*model.Flag{
			{Data: &model.Data{
				Name:        "petName",
				Description: "Generate random pet name"},
			},
			{Data: &model.Data{
				Name:        "cat",
				Description: "Generate random cat"},
			},
			{Data: &model.Data{
				Name:        "dog",
				Description: "Generate random dog"},
			},
			{Data: &model.Data{
				Name:        "bird",
				Description: "Generate random bird"},
			},
		},
	}
}

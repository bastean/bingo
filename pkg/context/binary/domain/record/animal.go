package record

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewAnimal() *model.Record {
	return &model.Record{
		Command: model.NewCommand("animal", "Generate random animal information"),
		Flags: []*model.Flag{
			{Name: "petName", Description: "Generate random pet name"},
			{Name: "cat", Description: "Generate random cat"},
			{Name: "dog", Description: "Generate random dog"},
			{Name: "bird", Description: "Generate random bird"},
		},
	}
}

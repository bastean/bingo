package aggregate

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/command"
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Root struct {
	*model.Command
}

func create(name, description string) *Root {
	nameVO := valueObject.NewName(name)
	descriptionVO := valueObject.NewDescription(description)

	return &Root{
		&model.Command{
			Data: &model.Data{
				Name:        nameVO.Value,
				Description: descriptionVO.Value,
				Switch:      true,
			},
			Subcommands: []*model.Command{
				command.NewPerson(),
				command.NewAddress(),
				command.NewPayment(),
				command.NewInternet(),
				command.NewLanguage(),
				command.NewAnimal(),
			},
		},
	}
}

func NewRoot(name, description string) *Root {
	return create(name, description)
}

package aggregate

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/domain/record"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Binary struct {
	*model.Record
}

func create(name, description string) *Binary {
	nameVO := valueObject.NewName(name)
	descriptionVO := valueObject.NewDescription(description)

	command := model.NewCommand(nameVO.Value, descriptionVO.Value)

	arguments := []*model.Argument{}

	options := []*model.Option{}

	flags := []*model.Flag{}

	subcommands := []*model.Record{
		record.NewPerson(),
		record.NewAddress(),
		record.NewPayment(),
		record.NewInternet(),
		record.NewLanguage(),
		record.NewAnimal(),
	}

	records := model.NewRecord(command, arguments, options, flags, subcommands)

	return &Binary{
		Record: records,
	}
}

func NewBinary(name, description string) *Binary {
	return create(name, description)
}

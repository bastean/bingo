package aggregate

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/command"
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

type Binary struct {
	*valueObject.Platform
	*valueObject.Filename
	*model.Command
}

func create(os, arch, filename, description string) *Binary {
	platformVO := valueObject.NewPlatform(os, arch)
	filenameVO := valueObject.NewFilename(filename)
	descriptionVO := valueObject.NewDescription(description)

	return &Binary{
		Platform: platformVO,
		Filename: filenameVO,
		Command: &model.Command{
			Data: &model.Data{
				Name:        filenameVO.Value,
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

func NewBinary(os, arch, filename, description string) *Binary {
	return create(os, arch, filename, description)
}

package command

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewPerson() *model.Command {
	return &model.Command{
		Data: &model.Data{
			Name:        "person",
			Description: "Generate random person information",
		},
		Flags: []*model.Flag{
			{Data: &model.Data{
				Name:        "firstName",
				Description: "Generate random first name"},
			},
			{Data: &model.Data{
				Name:        "lastName",
				Description: "Generate random last name"},
			},
			{Data: &model.Data{
				Name:        "ssn",
				Description: "Generate random ssn"},
			},
			{Data: &model.Data{
				Name:        "email",
				Description: "Generate random email"},
			},
			{Data: &model.Data{
				Name:        "phone",
				Description: "Generate random phone"},
			},
		},
	}
}

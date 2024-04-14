package record

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewPerson() *model.Record {
	return &model.Record{
		Command: model.NewCommand("person", "Generate random person information"),
		Flags: []*model.Flag{
			{Name: "firstName", Description: "Generate random first name"},
			{Name: "lastName", Description: "Generate random last name"},
			{Name: "ssn", Description: "Generate random ssn"},
			{Name: "email", Description: "Generate random email"},
			{Name: "phone", Description: "Generate random phone"},
		},
	}
}

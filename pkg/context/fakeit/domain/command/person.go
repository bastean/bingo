package command

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Person struct {
	*model.Command
}

func NewPerson() *Person {
	return &Person{
		&model.Command{
			Name:        "person",
			Description: "Generate random person information",
			Arguments: []*model.Argument{
				{Name: "firstName", Description: "Generate random first name"},
				{Name: "lastName", Description: "Generate random last name"},
				{Name: "ssn", Description: "Generate random ssn"},
				{Name: "email", Description: "Generate random email"},
				{Name: "phone", Description: "Generate random phone"},
			},
		},
	}
}

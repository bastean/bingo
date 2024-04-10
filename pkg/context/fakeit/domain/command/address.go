package command

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Address struct {
	*model.Command
}

func NewAddress() *Address {
	return &Address{
		&model.Command{
			Name:        "address",
			Description: "Generate random address information",
			Arguments: []*model.Argument{
				{Name: "country", Description: "Generate random country"},
				{Name: "state", Description: "Generate random state"},
				{Name: "city", Description: "Generate random city"},
				{Name: "street", Description: "Generate random street"},
				{Name: "zip", Description: "Generate random zip"},
			},
		},
	}
}

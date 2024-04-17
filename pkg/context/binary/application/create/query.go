package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

type Query struct {
	*model.Command
}

func NewQuery(name, description string) *Query {
	return &Query{
		Command: &model.Command{
			Data: &model.Data{
				Name:        name,
				Description: description,
				Switch:      true,
			},
		},
	}
}

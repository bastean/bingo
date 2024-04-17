package command

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewAddress() *model.Command {
	return &model.Command{
		Data: &model.Data{
			Name:        "address",
			Description: "Generate random address information",
		},
		Flags: []*model.Flag{
			{Data: &model.Data{
				Name:        "country",
				Description: "Generate random country"},
			},
			{Data: &model.Data{
				Name:        "state",
				Description: "Generate random state"},
			},
			{Data: &model.Data{
				Name:        "city",
				Description: "Generate random city"},
			},
			{Data: &model.Data{
				Name:        "street",
				Description: "Generate random street"},
			},
			{Data: &model.Data{
				Name:        "zip",
				Description: "Generate random zip"},
			},
		},
	}
}

package record

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewAddress() *model.Record {
	return &model.Record{
		Command: model.NewCommand("address", "Generate random address information"),
		Flags: []*model.Flag{
			{Name: "country", Description: "Generate random country"},
			{Name: "state", Description: "Generate random state"},
			{Name: "city", Description: "Generate random city"},
			{Name: "street", Description: "Generate random street"},
			{Name: "zip", Description: "Generate random zip"},
		},
	}
}

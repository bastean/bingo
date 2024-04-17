package command

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewInternet() *model.Command {
	return &model.Command{
		Data: &model.Data{
			Name:        "internet",
			Description: "Generate random internet information",
		},
		Flags: []*model.Flag{
			{Data: &model.Data{
				Name:        "url",
				Description: "Generate random url"},
			},
			{Data: &model.Data{
				Name:        "domainName",
				Description: "Generate random domain name"},
			},
			{Data: &model.Data{
				Name:        "domainSuffix",
				Description: "Generate random domain suffix"},
			},
			{Data: &model.Data{
				Name:        "macAddress",
				Description: "Generate random mac address"},
			},
			{Data: &model.Data{
				Name:        "userAgent",
				Description: "Generate random user agent"},
			},
		},
	}
}

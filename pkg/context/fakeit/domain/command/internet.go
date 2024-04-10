package command

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Internet struct {
	*model.Command
}

func NewInternet() *Internet {
	return &Internet{
		&model.Command{
			Name:        "internet",
			Description: "Generate random internet information",
			Arguments: []*model.Argument{
				{Name: "url", Description: "Generate random url"},
				{Name: "domainName", Description: "Generate random domain name"},
				{Name: "domainSuffix", Description: "Generate random domain suffix"},
				{Name: "macAddress", Description: "Generate random mac address"},
				{Name: "userAgent", Description: "Generate random user agent"},
			},
		},
	}
}

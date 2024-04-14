package record

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewInternet() *model.Record {
	return &model.Record{
		Command: model.NewCommand("internet", "Generate random internet information"),
		Flags: []*model.Flag{
			{Name: "url", Description: "Generate random url"},
			{Name: "domainName", Description: "Generate random domain name"},
			{Name: "domainSuffix", Description: "Generate random domain suffix"},
			{Name: "macAddress", Description: "Generate random mac address"},
			{Name: "userAgent", Description: "Generate random user agent"},
		},
	}
}

package command

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewPayment() *model.Command {
	return &model.Command{
		Data: &model.Data{
			Name:        "payment",
			Description: "Generate random payment information",
		},
		Flags: []*model.Flag{
			{Data: &model.Data{
				Name:        "creditCardNumber",
				Description: "Generate random credit card number"},
			},
			{Data: &model.Data{
				Name:        "creditCardCvv",
				Description: "Generate random credit card cvv"},
			},
			{Data: &model.Data{
				Name:        "creditCardExp",
				Description: "Generate random credit card exp"},
			},
			{Data: &model.Data{
				Name:        "bitcoinAddress",
				Description: "Generate random bitcoin address"},
			},
			{Data: &model.Data{
				Name:        "bitcoinPrivateKey",
				Description: "Generate random bitcoin private key"},
			},
		},
	}
}

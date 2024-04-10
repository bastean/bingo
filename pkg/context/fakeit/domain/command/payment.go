package command

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Payment struct {
	*model.Command
}

func NewPayment() *Payment {
	return &Payment{
		&model.Command{
			Name:        "payment",
			Description: "Generate random payment information",
			Arguments: []*model.Argument{
				{Name: "creditCardNumber", Description: "Generate random credit card number"},
				{Name: "creditCardCvv", Description: "Generate random credit card cvv"},
				{Name: "creditCardExp", Description: "Generate random credit card exp"},
				{Name: "bitcoinAddress", Description: "Generate random bitcoin address"},
				{Name: "bitcoinPrivateKey", Description: "Generate random bitcoin private key"},
			},
		},
	}
}

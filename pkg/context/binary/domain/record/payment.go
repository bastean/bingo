package record

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewPayment() *model.Record {
	return &model.Record{
		Command: model.NewCommand("payment", "Generate random payment information"),
		Flags: []*model.Flag{
			{Name: "creditCardNumber", Description: "Generate random credit card number"},
			{Name: "creditCardCvv", Description: "Generate random credit card cvv"},
			{Name: "creditCardExp", Description: "Generate random credit card exp"},
			{Name: "bitcoinAddress", Description: "Generate random bitcoin address"},
			{Name: "bitcoinPrivateKey", Description: "Generate random bitcoin private key"},
		},
	}
}

package payment

import (
	"fmt"

	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var PaymentFlags = []*config.Flag{
	{Name: "creditCardNumber", Run: func() {
		fmt.Printf("Credit Card Number: %s\n", service.Create.CreditCardNumber(nil))
	}},
	{Name: "creditCardCvv", Run: func() {
		fmt.Printf("Credit Card Cvv: %s\n", service.Create.CreditCardCvv())
	}},
	{Name: "creditCardExp", Run: func() {
		fmt.Printf("Credit Card Exp: %s\n", service.Create.CreditCardExp())
	}},
	{Name: "bitcoinAddress", Run: func() {
		fmt.Printf("Bitcoin Address: %s\n", service.Create.BitcoinAddress())
	}},
	{Name: "bitcoinPrivateKey", Run: func() {
		fmt.Printf("Bitcoin Private Key: %s\n", service.Create.BitcoinPrivateKey())
	}},
}

var PaymentCmd = &cobra.Command{
	Use: "payment",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunCommandFlags(cmd, PaymentFlags)
	},
}

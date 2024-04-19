package cmd

import (
	"fmt"
	"os"

	"github.com/bastean/bingo/cmd/address"
	"github.com/bastean/bingo/cmd/animal"
	"github.com/bastean/bingo/cmd/internet"
	"github.com/bastean/bingo/cmd/language"
	"github.com/bastean/bingo/cmd/payment"
	"github.com/bastean/bingo/cmd/person"
	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   config.Root.Name,
	Short: config.Root.Description,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			service.ShowHelp(cmd)
		}
	},
}

func init() {
	RootCmd.AddCommand(person.PersonCmd)
	RootCmd.AddCommand(address.AddressCmd)
	RootCmd.AddCommand(payment.PaymentCmd)
	RootCmd.AddCommand(internet.InternetCmd)
	RootCmd.AddCommand(language.LanguageCmd)
	RootCmd.AddCommand(animal.AnimalCmd)
}

func Execute() {
	service.Switcher(RootCmd, config.Root)

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

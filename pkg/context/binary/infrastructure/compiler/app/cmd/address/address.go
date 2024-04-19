package address

import (
	"fmt"

	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var AddressFlags = []*config.Flag{
	{Name: "country", Run: func() {
		fmt.Printf("Country: %s\n", service.Create.Country())
	}},
	{Name: "state", Run: func() {
		fmt.Printf("State: %s\n", service.Create.State())
	}},
	{Name: "city", Run: func() {
		fmt.Printf("City: %s\n", service.Create.City())
	}},
	{Name: "street", Run: func() {
		fmt.Printf("Street: %s\n", service.Create.Street())
	}},
	{Name: "zip", Run: func() {
		fmt.Printf("Zip: %s\n", service.Create.Zip())
	}},
}

var AddressCmd = &cobra.Command{
	Use: "address",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunCommandFlags(cmd, AddressFlags)
	},
}

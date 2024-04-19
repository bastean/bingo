package person

import (
	"fmt"

	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var PersonFlags = []*config.Flag{
	{Name: "firstName", Run: func() {
		fmt.Printf("First Name: %s\n", service.Create.FirstName())
	}},
	{Name: "lastName", Run: func() {
		fmt.Printf("Last Name: %s\n", service.Create.LastName())
	}},
	{Name: "ssn", Run: func() {
		fmt.Printf("SSN: %s\n", service.Create.SSN())
	}},
	{Name: "email", Run: func() {
		fmt.Printf("Email: %s\n", service.Create.Email())
	}},
	{Name: "phone", Run: func() {
		fmt.Printf("Phone: %s\n", service.Create.Phone())
	}},
}

var PersonCmd = &cobra.Command{
	Use: "person",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunCommandFlags(cmd, PersonFlags)
	},
}

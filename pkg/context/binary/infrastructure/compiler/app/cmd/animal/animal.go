package animal

import (
	"fmt"

	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var AnimalFlags = []*config.Flag{
	{Name: "petName", Run: func() {
		fmt.Printf("Pet Name: %s\n", service.Create.PetName())
	}},
	{Name: "cat", Run: func() {
		fmt.Printf("Cat: %s\n", service.Create.Cat())
	}},
	{Name: "dog", Run: func() {
		fmt.Printf("Dog: %s\n", service.Create.Dog())
	}},
	{Name: "bird", Run: func() {
		fmt.Printf("Bird: %s\n", service.Create.Bird())
	}},
}

var AnimalCmd = &cobra.Command{
	Use: "animal",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunCommandFlags(cmd, AnimalFlags)
	},
}

package language

import (
	"fmt"

	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var LanguageFlags = []*config.Flag{
	{Name: "programmingLanguage", Run: func() {
		fmt.Printf("Programming Language: %s\n", service.Create.ProgrammingLanguage())
	}},
	{Name: "bestProgrammingLanguage", Run: func() {
		fmt.Println("Best Programming Language: Go")
	}},
}

var LanguageCmd = &cobra.Command{
	Use: "language",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunCommandFlags(cmd, LanguageFlags)
	},
}

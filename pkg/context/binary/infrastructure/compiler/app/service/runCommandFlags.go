package service

import (
	"github.com/bastean/bingo/config"
	"github.com/spf13/cobra"
)

func RunCommandFlags(cmd *cobra.Command, flags []*config.Flag) {
	hasNoFlags := true

	for _, flag := range flags {
		isUsed, _ := cmd.Flags().GetBool(flag.Name)

		if isUsed {
			flag.Run()
			hasNoFlags = false
		}
	}

	if hasNoFlags {
		ShowHelp(cmd)
	}
}

package service

import (
	"os"

	"github.com/spf13/cobra"
)

func ShowHelp(cmd *cobra.Command) {
	cmd.Help()
	os.Exit(0)
}

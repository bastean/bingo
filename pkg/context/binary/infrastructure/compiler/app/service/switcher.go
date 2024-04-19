package service

import (
	"github.com/bastean/bingo/model"
	"github.com/spf13/cobra"
)

func Switcher(cmd *cobra.Command, config *model.Command) {
	if !config.Switch && cmd.HasParent() {
		cmd.Parent().RemoveCommand(cmd)
		return
	}

	cmd.Use = config.Name
	cmd.Short = config.Description

	// TODO?: cmd.Args

	for _, flag := range config.Flags {
		if flag.Switch {
			cmd.Flags().Bool(flag.Name, false, flag.Description)
		}
	}

	for _, subCmdConfig := range config.Subcommands {
		for _, subCmd := range cmd.Commands() {
			if subCmdConfig.Name == subCmd.Use {
				Switcher(subCmd, subCmdConfig)
			}
		}
	}
}

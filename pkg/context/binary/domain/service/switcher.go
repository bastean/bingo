package service

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func Switcher(model *model.Command, switches *model.Command) {
	if model.Name == switches.Name {
		model.Switch = switches.Switch
	}

	for _, modelArgument := range model.Arguments {
		for _, switchArgument := range switches.Arguments {
			if modelArgument.Name == switchArgument.Name {
				modelArgument.Switch = switchArgument.Switch
			}
		}
	}

	for _, modelFlag := range model.Flags {
		for _, switchFlag := range switches.Flags {
			if modelFlag.Name == switchFlag.Name {
				modelFlag.Switch = switchFlag.Switch
			}
		}
	}

	for _, modelSubcommand := range model.Subcommands {
		for _, switchSubcommand := range switches.Subcommands {
			if modelSubcommand.Name == switchSubcommand.Name {
				Switcher(modelSubcommand, switchSubcommand)
			}
		}
	}
}

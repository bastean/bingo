package model

type Command struct {
	*Data
	Arguments   []*Argument
	Flags       []*Flag
	Subcommands []*Command
}

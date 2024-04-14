package model

type Record struct {
	*Command
	Arguments   []*Argument
	Options     []*Option
	Flags       []*Flag
	Subcommands []*Record
}

func NewRecord(command *Command, arguments []*Argument, options []*Option, flags []*Flag, subcommands []*Record) *Record {
	return &Record{
		Command:     command,
		Arguments:   arguments,
		Options:     options,
		Flags:       flags,
		Subcommands: subcommands,
	}
}

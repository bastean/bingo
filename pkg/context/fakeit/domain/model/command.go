package model

type Command struct {
	Name        string
	Description string
	Arguments   []*Argument
}

func NewCommand(name, description string, arguments []*Argument) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Arguments:   arguments,
	}
}

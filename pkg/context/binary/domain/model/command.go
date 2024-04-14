package model

type Command struct {
	Name        string
	Description string
}

func NewCommand(name, description string) *Command {
	return &Command{
		Name:        name,
		Description: description,
	}
}

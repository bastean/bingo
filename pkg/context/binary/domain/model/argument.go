package model

type Argument struct {
	Name        string
	Description string
}

func NewArgument(name, description string) *Argument {
	return &Argument{
		Name:        name,
		Description: description,
	}
}

package model

type Option struct {
	Name        string
	Description string
}

func NewOption(name, description string) *Argument {
	return &Argument{
		Name:        name,
		Description: description,
	}
}

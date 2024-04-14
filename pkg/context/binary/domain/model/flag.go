package model

type Flag struct {
	Name        string
	Description string
}

func NewFlag(name, description string) *Flag {
	return &Flag{
		Name:        name,
		Description: description,
	}
}

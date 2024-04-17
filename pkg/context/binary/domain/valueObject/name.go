package valueObject

type Name struct {
	Value string
}

func NewName(name string) *Name {
	if name == "" {
		name = "bingo"
	}

	return &Name{
		Value: name,
	}
}

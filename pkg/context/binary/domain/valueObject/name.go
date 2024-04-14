package valueObject

type Name struct {
	Value string
}

func NewName(name string) *Name {
	if name == "" {
		name = "binary"
	}

	return &Name{
		Value: name,
	}
}

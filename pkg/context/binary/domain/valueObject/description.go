package valueObject

type Description struct {
	Value string
}

func NewDescription(description string) *Description {
	if description == "" {
		description = "Custom executable binary"
	}

	return &Description{
		Value: description,
	}
}

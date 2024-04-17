package valueObject

type Description struct {
	Value string
}

func NewDescription(description string) *Description {
	if description == "" {
		description = "Example service for generating cross-platform custom executable binary"
	}

	return &Description{
		Value: description,
	}
}

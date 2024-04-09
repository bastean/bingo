package valueObject

import (
	sharedValueObject "github.com/bastean/bingo/pkg/context/shared/domain/valueObject"
)

type Email struct {
	Value string
}

func NewEmail(email string) *Email {
	return &Email{
		Value: sharedValueObject.NewEmail(email).Value,
	}
}

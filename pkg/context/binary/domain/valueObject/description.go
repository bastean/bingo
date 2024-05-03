package valueObject

import (
	"regexp"
	"strings"

	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
)

const DescriptionMinCharactersLength = "1"
const DescriptionMaxCharactersLength = "100"

var InvalidDescriptionValue = errors.InvalidValue{Message: "Description must be between " + DescriptionMinCharactersLength + " to " + DescriptionMaxCharactersLength + " characters"}

type Description struct {
	Value string
}

func ensureIsValidDescription(description string) {
	validate := regexp.MustCompile(`^[\w\s.:-]{1,100}$`)

	if !validate.MatchString(description) {
		panic(InvalidDescriptionValue)
	}
}

func NewDescription(description string) *Description {
	if description == "" {
		description = "Example service for generating custom cross-compile executable binaries"
	}

	description = strings.TrimSpace(description)

	ensureIsValidDescription(description)

	descriptionVO := &Description{description}

	return descriptionVO
}

package argument

import (
	"fmt"

	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Usage struct {
	*model.Argument
}

func NewUsage(name string) *Usage {
	return &Usage{
		Argument: &model.Argument{
			Name:        name,
			Description: fmt.Sprintf("%s <command> [arguments]", name),
		},
	}
}

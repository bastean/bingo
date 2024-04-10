package command

import (
	"github.com/bastean/bingo/pkg/context/fakeit/domain/model"
)

type Language struct {
	*model.Command
}

func NewLanguage() *Language {
	return &Language{
		&model.Command{
			Name:        "language",
			Description: "Generate random language information",
			Arguments: []*model.Argument{
				{Name: "programmingLanguage", Description: "Generate random programming language"},
				{Name: "programmingLanguageBest", Description: "Generate random programming languageBest"},
			},
		},
	}
}

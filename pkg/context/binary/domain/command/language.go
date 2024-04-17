package command

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewLanguage() *model.Command {
	return &model.Command{
		Data: &model.Data{
			Name:        "language",
			Description: "Generate random language information",
		},
		Flags: []*model.Flag{
			{Data: &model.Data{
				Name:        "programmingLanguage",
				Description: "Generate random programming language"},
			},
			{Data: &model.Data{
				Name:        "programmingLanguageBest",
				Description: "Generate random programming languageBest"},
			},
		},
	}
}

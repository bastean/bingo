package record

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

func NewLanguage() *model.Record {
	return &model.Record{
		Command: model.NewCommand("language", "Generate random language information"),
		Flags: []*model.Flag{
			{Name: "programmingLanguage", Description: "Generate random programming language"},
			{Name: "programmingLanguageBest", Description: "Generate random programming languageBest"},
		},
	}
}

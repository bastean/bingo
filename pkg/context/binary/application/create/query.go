package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

type Platform struct {
	OS   string
	Arch string
}

type Query struct {
	Platform *Platform
	Filename string
	*model.Command
}

func NewQuery(os, arch, filename, description string) *Query {
	return &Query{
		Platform: &Platform{
			OS:   os,
			Arch: arch,
		},
		Filename: filename,
		Command: &model.Command{
			Data: &model.Data{
				Name:        filename,
				Description: description,
				Switch:      true,
			},
		},
	}
}

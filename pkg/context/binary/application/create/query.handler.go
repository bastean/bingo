package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/model"
)

type QueryHandler struct {
	*Create
}

func (handler *QueryHandler) customBinaryRecord(binaryModel *Record, binary *model.Record) {
	if !binaryModel.Command.ShouldRecorded {
		binary.Command = nil
		// TODO?(cleanest struct): return
	}

	for i, argument := range binary.Arguments {
		for _, argumentModel := range binaryModel.Arguments {
			if argument.Name == argumentModel.Name && !argumentModel.ShouldRecorded {
				binary.Arguments[i] = nil
			}
		}
	}

	for i, option := range binary.Options {
		for _, optionModel := range binaryModel.Options {
			if option.Name == optionModel.Name && !optionModel.ShouldRecorded {
				binary.Options[i] = nil
			}
		}
	}

	for i, flag := range binary.Flags {
		for _, flagModel := range binaryModel.Flags {
			if flag.Name == flagModel.Name && !flagModel.ShouldRecorded {
				binary.Flags[i] = nil
			}
		}
	}

	for _, subcommandModel := range binaryModel.Subcommands {
		for _, subcommand := range binary.Subcommands {
			switch {
			case subcommand.Command == nil:
				continue
			case subcommand.Command.Name == subcommandModel.Command.Name:
				handler.customBinaryRecord(subcommandModel, subcommand)
			}
		}
	}
}

func (handler *QueryHandler) Handle(query *Query) *Response {
	binary := aggregate.NewBinary(query.Name, query.Description)

	binaryModel := query.Record

	binaryModel.Command.ShouldRecorded = true

	handler.customBinaryRecord(binaryModel, binary.Record)

	filePath := handler.Create.Run(binary)

	return NewResponse(filePath.Value)
}

func NewQueryHandler(create *Create) *QueryHandler {
	return &QueryHandler{
		Create: create,
	}
}

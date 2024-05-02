package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/service"
)

type QueryHandler struct {
	*Create
}

func (handler *QueryHandler) Handle(query *Query) *Response {
	query.Command.Name = query.Filename

	query.Command.Switch = true

	binary := aggregate.NewBinary(query.Platform.OS, query.Platform.Arch, query.Filename, query.Command.Description)

	model := binary.Command

	switches := query.Command

	service.Switcher(model, switches)

	filePath := handler.Create.Run(binary)

	return NewResponse(filePath.Value)
}

func NewQueryHandler(create *Create) *QueryHandler {
	return &QueryHandler{
		Create: create,
	}
}

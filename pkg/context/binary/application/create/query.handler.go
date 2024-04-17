package create

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/service"
)

type QueryHandler struct {
	*Create
}

func (handler *QueryHandler) Handle(query *Query) *Response {
	query.Command.Switch = true

	root := aggregate.NewRoot(query.Command.Name, query.Command.Description)

	model := root.Command

	switches := query.Command

	service.Switcher(model, switches)

	filePath := handler.Create.Run(root)

	return NewResponse(filePath.Value)
}

func NewQueryHandler(create *Create) *QueryHandler {
	return &QueryHandler{
		Create: create,
	}
}

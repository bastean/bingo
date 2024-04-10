package builder

type QueryHandler struct {
	*Builder
}

func (handler *QueryHandler) Handle(query *Query) *Response {
	// TODO: fakeit builder query and response handler

	return NewResponse("success")
}

func NewQueryHandler(builder *Builder) *QueryHandler {
	return &QueryHandler{
		Builder: builder,
	}
}

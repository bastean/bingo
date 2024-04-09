package verify

import "github.com/bastean/bingo/pkg/context/user/domain/valueObject"

type CommandHandler struct {
	*Verify
}

func (handler *CommandHandler) Handle(command *Command) {
	idVO := valueObject.NewId(command.Id)

	handler.Verify.Run(idVO)
}

func NewCommandHandler(verify *Verify) *CommandHandler {
	return &CommandHandler{
		Verify: verify,
	}
}

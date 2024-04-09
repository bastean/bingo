package communicationMock

import (
	"github.com/bastean/bingo/pkg/context/notify/domain/model"
	"github.com/stretchr/testify/mock"
)

type MailMock struct {
	mock.Mock
}

func (m *MailMock) Send(template model.MailTemplate) {
	m.Called(template)
}

func NewMailMock() *MailMock {
	return new(MailMock)
}

package model

import (
	"github.com/bastean/bingo/pkg/context/shared/domain/exchange"
	"github.com/bastean/bingo/pkg/context/shared/domain/message"
	"github.com/bastean/bingo/pkg/context/shared/domain/queue"
)

type Broker interface {
	AddExchange(exchange *exchange.Exchange)
	AddQueue(queue *queue.Queue)
	AddQueueMessageBind(queue *queue.Queue, bindingKeys []string)
	AddQueueConsumer(consumer Consumer)
	PublishMessages(messages []*message.Message)
}

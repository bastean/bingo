package service

import (
	"github.com/bastean/bingo/pkg/context/shared/domain/queue"
)

func HasNoQueue(queues []*queue.Queue, queue *queue.Queue) bool {
	isNotPresent := true

	for _, present := range queues {
		if present.Name == queue.Name {
			isNotPresent = false
			break
		}
	}

	return isNotPresent
}

package messagebus

import (
	"sync"
)

type MessageBus struct {
	lock     sync.Mutex
	callback map[string][]chan interface{}
}

func New() *MessageBus {
	return &MessageBus{callbacks: make(map[string][]chan interface{})}
}

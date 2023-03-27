package messagebus

import (
	"sync"
)

type MessageBus struct {
	lock     sync.Mutex
	callBack map[string][]chan interface{}
}

func New() *MessageBus {
	return &MessageBus{callBack: make(map[string][]chan interface{})}
}

func (mb *MessageBus) Subscribe(topic string) <-chan interface{} {
	mb.lock.Lock()
	defer mb.lock.Unlock()

	ch := make(chan interface{}, 1)
	mb.callBack[topic] = append(mb.callBack[topic], ch)

	return ch
}

package messagebus

import (
	"sync"
)

type MessageBus struct {
	lock     sync.Mutex
	callback map[string][]chan interface{}
}

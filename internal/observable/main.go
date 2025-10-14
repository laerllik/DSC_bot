package main

import "sync"

type Event struct {
	Type string
	Data interface{}
}

type Observer struct {
	subscribers map[int]func(e Event)
	lastAddedId int
	mu          sync.Mutex
}

func (o *Observer) AddSubscriber(e string, f func(e Event)) int {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.subscribers == nil {
		o.subscribers = map[int]func(e Event){}
	}

	o.lastAddedId = o.lastAddedId + 1
	o.subscribers[o.lastAddedId-1] = f

	return o.lastAddedId - 1
}

func (o *Observer) RemoveSubscriber(id int) {
	o.mu.Lock()
	defer o.mu.Unlock()

	delete(o.subscribers, id)
}

func (o *Observer) EmitEvent(e Event) {
	o.mu.Lock()
	defer o.mu.Unlock()

	for _, f := range o.subscribers {
		f(e)
	}
}

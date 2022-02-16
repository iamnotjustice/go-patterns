package observer

import (
	"log"

	"github.com/pborman/uuid"
)

type Subscriber interface {
	OnNotify(Event)
}

type ConcreteSubscriber struct {
	id string

	GotEvent bool
}

func NewSubscriber() *ConcreteSubscriber {
	return &ConcreteSubscriber{
		id: uuid.New(),
	}
}

func (o *ConcreteSubscriber) OnNotify(e Event) {
	// Do stuff when we got an event
	o.GotEvent = true
	log.Printf("*** Subscriber %s received: %d\n", o.id, e.Data)
}

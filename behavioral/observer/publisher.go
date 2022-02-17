package observer

type Event struct {
	Data int64
}

type Publisher interface {
	Register(Subscriber)

	Deregister(Subscriber)

	Notify(Event)
}

type ConcretePublisher struct {
	subscribers map[Subscriber]struct{}
}

func NewPublisher() Publisher {
	return &ConcretePublisher{
		subscribers: map[Subscriber]struct{}{},
	}
}

func (o *ConcretePublisher) Register(l Subscriber) {
	o.subscribers[l] = struct{}{}
}

func (o *ConcretePublisher) Deregister(l Subscriber) {
	delete(o.subscribers, l)
}

func (p *ConcretePublisher) Notify(e Event) {
	for o := range p.subscribers {
		o.OnNotify(e)
	}
}

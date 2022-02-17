package observer_test

import (
	"testing"
	"time"

	"github.com/iamnotjustice/go-patterns/behavioral/observer"
	"github.com/stretchr/testify/assert"
)

func Test_Observer(t *testing.T) {
	// Initialize a new Publisher.
	n := observer.NewPublisher()

	// Register a couple of subscribers.
	n.Register(observer.NewSubscriber())
	n.Register(observer.NewSubscriber())

	sub1 := observer.NewSubscriber()
	sub2 := observer.NewSubscriber()

	n.Register(sub1)
	n.Register(sub2)

	n.Notify(observer.Event{Data: time.Now().UnixNano()})

	assert.Equal(t, true, sub1.GotEvent)
	assert.Equal(t, true, sub2.GotEvent)
}

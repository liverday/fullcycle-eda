package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetTimestamp() time.Time
	GetPayload() interface{}
}

type EventHandler interface {
	Handle(data Event, wg *sync.WaitGroup)
}

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Dispatch(eventName string, event Event) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	FindIndex(eventName string, handler EventHandler) int
	Clear() error
}

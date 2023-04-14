package events

import (
	"errors"
	"sync"
)

type EventDispatcherImpl struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() *EventDispatcherImpl {
	return &EventDispatcherImpl{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *EventDispatcherImpl) Has(eventName string, handler EventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcherImpl) Dispatch(event Event) {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}

		wg.Wait()
	}
}

func (ed *EventDispatcherImpl) Register(eventName string, handler EventHandler) error {
	if ok := ed.Has(eventName, handler); ok {
		return errors.New("handler already registered")
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcherImpl) Clear() error {
	ed.handlers = make(map[string][]EventHandler)
	return nil
}

func (ed *EventDispatcherImpl) FindIndex(eventName string, handler EventHandler) int {
	if handlers, ok := ed.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				return i
			}
		}
	}

	return -1
}

func (ed *EventDispatcherImpl) Remove(eventName string, handler EventHandler) error {
	idx := ed.FindIndex(eventName, handler)

	if idx < 0 {
		return errors.New("handler not found")
	}

	ed.handlers[eventName] = append(ed.handlers[eventName][:idx], ed.handlers[eventName][idx+1:]...)

	return nil
}

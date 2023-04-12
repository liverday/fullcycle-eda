package events

import "errors"

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

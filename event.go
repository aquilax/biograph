package biograph

import "time"

// EventType contains the type of the event
type EventType string

// LifeEvent is a single life event interface
type LifeEvent interface {
	getType() EventType
	getFrom() time.Time
	getTo() time.Time
	getName() string
	getMeta() map[string]string
}

type Event struct {
	eventType EventType
	from      time.Time
	to        time.Time
	name      string
	meta      map[string]string
}

func NewEvent() *Event {
	return &Event{}
}

func (e *Event) getType() EventType         { return e.eventType }
func (e *Event) getFrom() time.Time         { return e.from }
func (e *Event) getTo() time.Time           { return e.to }
func (e *Event) getName() string            { return e.name }
func (e *Event) getMeta() map[string]string { return e.meta }

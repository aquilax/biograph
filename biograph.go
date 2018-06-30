package biograph

import "time"

// EventType contains the type of the event
type EventType string

// Event is a single life event
type Event struct {
	EventType EventType
	From      time.Time
	To        time.Time
	Name      string
	Meta      map[string]string
}

// LifeEvents handles all life events
type LifeEvents interface {
	Add(event Event) error
}

// Life handler
type Life struct {
	from   time.Time
	to     time.Time
	events []Event
}

// NewLife creates new Life handler
func NewLife(from, to time.Time) *Life {
	return &Life{from, to, []Event{}}
}

// Add new event to Life
func (l *Life) Add(event Event) error {
	// TODO: check if event has valid time boundaries
	l.events = append(l.events, event)
	return nil
}

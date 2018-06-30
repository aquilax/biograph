package biograph

import "time"

// LifeEvents handles all life events
type LifeEvents interface {
	Add(event LifeEvent) error
	Count() int
}

// Life handler
type Life struct {
	from   time.Time
	to     time.Time
	events []LifeEvent
}

// NewLife creates new Life handler
func NewLife(from, to time.Time) *Life {
	return &Life{from, to, []LifeEvent{}}
}

// Add new event to Life
func (l *Life) Add(event LifeEvent) error {
	// TODO: check if event has valid time boundaries
	l.events = append(l.events, event)
	return nil
}

// Count returns number of life events
func (l *Life) Count() int {
	return len(l.events)
}

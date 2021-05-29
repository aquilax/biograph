package biograph

import (
	"time"
)

// Life handles all life events
type Life interface {
	Add(event LifeEvent) error
	Items() Events
	Count() int
	From() time.Time
	To() time.Time
}

// Life handler
type LifeArray struct {
	from   time.Time
	to     time.Time
	events Events
}

// NewLife creates new Life handler
func NewLifeArray(from, to time.Time) *LifeArray {
	return &LifeArray{from, to, []LifeEvent{}}
}

// Add new event to Life
func (l *LifeArray) Add(events ...LifeEvent) error {
	l.events = append(l.events, events...)
	return nil
}

// Count returns number of life events
func (l *LifeArray) Count() int {
	return len(l.events)
}

// Items return all Events as an array
func (l *LifeArray) Items() Events {
	return l.events
}

// From returns the starting date
func (l *LifeArray) From() time.Time {
	return l.from
}

// To returns the end date
func (l *LifeArray) To() time.Time {
	return l.to
}

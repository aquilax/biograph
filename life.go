package biograph

import (
	"time"
)

// LifeEvents handles all life events
type Life interface {
	Add(event LifeEvent) error
	Items() Events
	Count() int
}

// Life handler
type LifeArray struct {
	from   time.Time
	to     time.Time
	events Events
}

// NewLife creates new Life handler
func NewLife(from, to time.Time) *LifeArray {
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

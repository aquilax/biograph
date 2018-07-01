package biograph

import (
	"sort"
	"time"
)

// LifeEvents handles all life events
type LifeEvents interface {
	Add(event LifeEvent) error
	Items() []LifeEvent
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
func (l *Life) Add(events ...LifeEvent) error {
	l.events = append(l.events, events...)
	return nil
}

// Count returns number of life events
func (l *Life) Count() int {
	return len(l.events)
}

func (l *Life) Items() []LifeEvent {
	return l.events
}

func (l *Life) Asc() []LifeEvent {
	sort.Slice(l.events, func(i, j int) bool { return l.events[i].getFrom().Before(l.events[j].getFrom()) })
	return l.events
}

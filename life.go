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
	Filter(f Filterer) []LifeEvent
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
	sort.Slice(l.events, func(i, j int) bool { return l.events[i].GetFrom().Before(l.events[j].GetFrom()) })
	return l.events
}

func (l *Life) Desc() []LifeEvent {
	sort.Slice(l.events, func(i, j int) bool { return l.events[j].GetFrom().Before(l.events[i].GetFrom()) })
	return l.events
}

func (l *Life) Filter(f Filterer) []LifeEvent {
	result := make([]LifeEvent, 0)
	for _, e := range l.events {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

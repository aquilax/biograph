package biograph

import "time"

type Filterer func(LifeEvent) bool

func NewBetween(from, to time.Time) Filterer {
	return func(l LifeEvent) bool {
		if l.GetTo().IsZero() {
			return from.Sub(l.GetFrom()) >= 0
		}
		if l.GetFrom().IsZero() {
			return l.GetTo().Sub(to) >= 0
		}
		return from.Sub(l.GetFrom()) >= 0 && l.GetTo().Sub(to) >= 0
	}
}

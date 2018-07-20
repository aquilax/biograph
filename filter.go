package biograph

import "time"

type Filterer func(LifeEvent) bool

func NewBetween(from, to time.Time) Filterer {
	return func(l LifeEvent) bool {
		return from.Sub(l.GetFrom()) >= 0 && l.GetTo().Sub(to) >= 0
	}
}

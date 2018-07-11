package biograph

import "time"

// EventType contains the type of the event
type EventType string

const (
	Home      EventType = "home"
	Education EventType = "education"
	Work      EventType = "work"
	Travel    EventType = "travel"
	Item      EventType = "item"
)

// LifeEvent is a single life event interface
type LifeEvent interface {
	getType() EventType
	getFrom() time.Time
	getTo() time.Time
	getName() string
	getMeta() *MetaData
}

type GenericEvent struct {
	from time.Time
	to   time.Time
	meta *MetaData
}

// HomeEvent represents home address change
type HomeEvent GenericEvent

// NewHome creates new Place to live
func NewHome(address, country string, from, to time.Time, meta *MetaData) *HomeEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"address": address, "country": country})
	return &HomeEvent{from, to, meta}
}

func (e *HomeEvent) getType() EventType { return Home }
func (e *HomeEvent) getFrom() time.Time { return e.from }
func (e *HomeEvent) getTo() time.Time   { return e.to }
func (e *HomeEvent) getName() string    { return e.meta.get("address") }
func (e *HomeEvent) getMeta() *MetaData { return e.meta }

// EducationEvent represents studying in educational institution
type EducationEvent GenericEvent

func NewEducation(school, degree string, from, to time.Time, meta *MetaData) *EducationEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"school": school, "degree": degree})
	return &EducationEvent{from, to, meta}
}

func (e *EducationEvent) getType() EventType { return Education }
func (e *EducationEvent) getFrom() time.Time { return e.from }
func (e *EducationEvent) getTo() time.Time   { return e.to }
func (e *EducationEvent) getName() string    { return e.meta.get("school") }
func (e *EducationEvent) getMeta() *MetaData { return e.meta }

type WorkEvent GenericEvent

func NewWork(employer, position string, from, to time.Time, meta *MetaData) *WorkEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"employer": employer, "position": position})
	return &WorkEvent{from, to, meta}
}

func (e *WorkEvent) getType() EventType { return Education }
func (e *WorkEvent) getFrom() time.Time { return e.from }
func (e *WorkEvent) getTo() time.Time   { return e.to }
func (e *WorkEvent) getName() string    { return e.meta.get("employer") }
func (e *WorkEvent) getMeta() *MetaData { return e.meta }

type TravelEvent GenericEvent

func NewTravel(place, country string, from, to time.Time, meta *MetaData) *TravelEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"place": place, "country": country})
	return &TravelEvent{from, to, meta}
}

func (e *TravelEvent) getType() EventType { return Travel }
func (e *TravelEvent) getFrom() time.Time { return e.from }
func (e *TravelEvent) getTo() time.Time   { return e.to }
func (e *TravelEvent) getName() string    { return e.meta.get("place") + ", " + e.meta.get("country") }
func (e *TravelEvent) getMeta() *MetaData { return e.meta }

type ItemEvent GenericEvent

func NewItem(category string, from, to time.Time, meta *MetaData) *ItemEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"category": category})
	return &ItemEvent{from, to, meta}
}

func (e *ItemEvent) getType() EventType { return Item }
func (e *ItemEvent) getFrom() time.Time { return e.from }
func (e *ItemEvent) getTo() time.Time   { return e.to }
func (e *ItemEvent) getName() string    { return e.meta.get("place") + ", " + e.meta.get("country") }
func (e *ItemEvent) getMeta() *MetaData { return e.meta }

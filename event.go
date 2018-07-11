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
	GetType() EventType
	GetFrom() time.Time
	GetTo() time.Time
	GetName() string
	GetMeta() *MetaData
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

func (e *HomeEvent) GetType() EventType { return Home }
func (e *HomeEvent) GetFrom() time.Time { return e.from }
func (e *HomeEvent) GetTo() time.Time   { return e.to }
func (e *HomeEvent) GetName() string    { return e.meta.get("address") }
func (e *HomeEvent) GetMeta() *MetaData { return e.meta }

// EducationEvent represents studying in educational institution
type EducationEvent GenericEvent

func NewEducation(school, degree string, from, to time.Time, meta *MetaData) *EducationEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"school": school, "degree": degree})
	return &EducationEvent{from, to, meta}
}

func (e *EducationEvent) GetType() EventType { return Education }
func (e *EducationEvent) GetFrom() time.Time { return e.from }
func (e *EducationEvent) GetTo() time.Time   { return e.to }
func (e *EducationEvent) GetName() string    { return e.meta.get("school") }
func (e *EducationEvent) GetMeta() *MetaData { return e.meta }

type WorkEvent GenericEvent

func NewWork(employer, position string, from, to time.Time, meta *MetaData) *WorkEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"employer": employer, "position": position})
	return &WorkEvent{from, to, meta}
}

func (e *WorkEvent) GetType() EventType { return Education }
func (e *WorkEvent) GetFrom() time.Time { return e.from }
func (e *WorkEvent) GetTo() time.Time   { return e.to }
func (e *WorkEvent) GetName() string    { return e.meta.get("employer") }
func (e *WorkEvent) GetMeta() *MetaData { return e.meta }

type TravelEvent GenericEvent

func NewTravel(place, country string, from, to time.Time, meta *MetaData) *TravelEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"place": place, "country": country})
	return &TravelEvent{from, to, meta}
}

func (e *TravelEvent) GetType() EventType { return Travel }
func (e *TravelEvent) GetFrom() time.Time { return e.from }
func (e *TravelEvent) GetTo() time.Time   { return e.to }
func (e *TravelEvent) GetName() string    { return e.meta.get("place") + ", " + e.meta.get("country") }
func (e *TravelEvent) GetMeta() *MetaData { return e.meta }

type ItemEvent GenericEvent

func NewItem(category string, from, to time.Time, meta *MetaData) *ItemEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"category": category})
	return &ItemEvent{from, to, meta}
}

func (e *ItemEvent) GetType() EventType { return Item }
func (e *ItemEvent) GetFrom() time.Time { return e.from }
func (e *ItemEvent) GetTo() time.Time   { return e.to }
func (e *ItemEvent) GetName() string    { return e.meta.get("place") + ", " + e.meta.get("country") }
func (e *ItemEvent) GetMeta() *MetaData { return e.meta }

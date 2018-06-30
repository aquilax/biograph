package biograph

import "time"

// EventType contains the type of the event
type EventType string

const (
	Home      EventType = "home"
	Education EventType = "education"
	Travel    EventType = "travel"
)

// LifeEvent is a single life event interface
type LifeEvent interface {
	getType() EventType
	getFrom() time.Time
	getTo() time.Time
	getName() string
	getMeta() *MetaData
}

// HomeEvent represents home address change
type HomeEvent struct {
	from time.Time
	to   time.Time
	meta *MetaData
}

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
type EducationEvent struct {
	from time.Time
	to   time.Time
	meta *MetaData
}

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

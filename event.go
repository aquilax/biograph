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
	Partner   EventType = "partner"
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

func (e *GenericEvent) GetFrom() time.Time { return e.from }
func (e *GenericEvent) GetTo() time.Time   { return e.to }
func (e *GenericEvent) GetMeta() *MetaData { return e.meta }

// HomeEvent represents home address change
type HomeEvent struct {
	*GenericEvent
}

// NewHome creates new Place to live
func NewHome(address, country string, from, to time.Time, meta *MetaData) *HomeEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"address": address, "country": country})
	return &HomeEvent{&GenericEvent{from, to, meta}}
}

func (e *HomeEvent) GetType() EventType { return Home }
func (e *HomeEvent) GetName() string    { return e.meta.Get("address") }

// EducationEvent represents studying in educational institution
type EducationEvent struct {
	*GenericEvent
}

func NewEducation(school, degree string, from, to time.Time, meta *MetaData) *EducationEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"school": school, "degree": degree})
	return &EducationEvent{&GenericEvent{from, to, meta}}
}

func (e *EducationEvent) GetType() EventType { return Education }
func (e *EducationEvent) GetName() string    { return e.meta.Get("school") }

type WorkEvent struct {
	*GenericEvent
}

func NewWork(employer, position string, from, to time.Time, meta *MetaData) *WorkEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"employer": employer, "position": position})
	return &WorkEvent{&GenericEvent{from, to, meta}}
}

func (e *WorkEvent) GetType() EventType { return Work }
func (e *WorkEvent) GetName() string    { return e.meta.Get("employer") }

type TravelEvent struct {
	*GenericEvent
}

func NewTravel(place, country string, from, to time.Time, meta *MetaData) *TravelEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"place": place, "country": country})
	return &TravelEvent{&GenericEvent{from, to, meta}}
}

func (e *TravelEvent) GetType() EventType { return Travel }
func (e *TravelEvent) GetName() string    { return e.meta.Get("place") + ", " + e.meta.Get("country") }

type ItemEvent struct {
	*GenericEvent
}

func NewItem(category string, from, to time.Time, meta *MetaData) *ItemEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"category": category})
	return &ItemEvent{&GenericEvent{from, to, meta}}
}

func (e *ItemEvent) GetType() EventType { return Item }
func (e *ItemEvent) GetName() string    { return e.meta.Get("category") }

type PartnerEvent struct {
	*GenericEvent
}

func NewPartner(name string, from, to time.Time, meta *MetaData) *PartnerEvent {
	if meta == nil {
		meta = &MetaData{}
	}
	meta.merge(&MetaData{"name": name})
	return &PartnerEvent{&GenericEvent{from, to, meta}}
}

func (e *PartnerEvent) GetType() EventType { return Partner }
func (e *PartnerEvent) GetName() string    { return e.meta.Get("name") }

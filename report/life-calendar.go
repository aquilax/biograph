package report

import (
	"encoding/json"
	"io"

	"github.com/aquilax/biograph"
)

const layoutISO = "2006-01-02"

type event struct {
	Type  int    `json:"type"`
	Date  string `json:"date"`
	Title string `json:"title"`
}

type lifeCalendar struct {
	Options map[string]string `json:"options"`
	Events  []event           `json:"events"`
}

func toInt(e biograph.EventType) int {
	switch e {
	case biograph.Home:
		return 1
	case biograph.Education:
		return 2
	case biograph.Work:
		return 3
	case biograph.Travel:
		return 4
	case biograph.Item:
		return 5
	case biograph.Partner:
		return 6
	case biograph.Roommate:
		return 7
	case biograph.Project:
		return 8
	case biograph.Document:
		return 9
	case biograph.Process:
		return 10
	}
	return -1
}

type LifeCalendar struct {
	out io.Writer
}

// NewLifeCalendar creates new report
func NewLifeCalendar(out io.Writer) *LifeCalendar {
	return &LifeCalendar{out}
}

// Generate generates life-calendar export
// https://github.com/ngduc/life-calendar
func (l *LifeCalendar) Generate(events biograph.Events) error {
	var lc lifeCalendar
	var e event
	var err error
	for _, event := range events {
		e.Type = toInt(event.GetType())
		e.Date = event.GetFrom().Format(layoutISO)
		e.Title = event.GetName()
		lc.Events = append(lc.Events, e)
	}
	bt, err := json.Marshal(lc)
	if err != nil {
		return err
	}
	_, err = l.out.Write(bt)
	return err
}

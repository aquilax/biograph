package report

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/aquilax/biograph"
)

const layoutISO = "2006-01-02"

type event struct {
	Type  int
	Date  string
	Title string
}

type lifeCalendar struct {
	Options map[string]string
	Events  []event
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
	out io.WriteCloser
}

// NewLifeCalendar creates new report
func NewLifeCalendar(out io.WriteCloser) *LifeCalendar {
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
	var b bytes.Buffer
	b.Write(bt)
	_, err = b.WriteTo(l.out)
	return err
}

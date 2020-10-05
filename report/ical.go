package report

import (
	"fmt"
	"io"

	"github.com/aquilax/biograph"
	"github.com/emersion/go-ical"
)

type ICal struct {
	out io.Writer
}

// NewICal creates new iCal reporter
func NewICal(out io.Writer) *ICal {
	return &ICal{out}
}

func (ic *ICal) Generate(events biograph.Events) error {
	cal := ical.NewCalendar()
	cal.Props.SetText(ical.PropVersion, "2.0")
	cal.Props.SetText(ical.PropProductID, "-//Biograph//NONSGML Calendar//EN")

	for _, e := range events {
		event := ical.NewEvent()
		event.Props.SetText(ical.PropUID, "uid@example.org")
		event.Props.SetText(ical.PropSummary, e.GetName())
		event.Props.SetText(ical.PropCategories, fmt.Sprintf("%v", e.GetType()))
		event.Props.SetDateTime(ical.PropDateTimeStamp, e.GetFrom())
		event.Props.SetDateTime(ical.PropDateTimeStart, e.GetTo())
		cal.Children = append(cal.Children, event.Component)
	}

	if err := ical.NewEncoder(ic.out).Encode(cal); err != nil {
		return err
	}
	return nil
}

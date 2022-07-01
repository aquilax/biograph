package report

import (
	"fmt"
	"io"
	"time"

	"github.com/aquilax/biograph"
)

type MarkWhen struct {
	out io.Writer
}

func NewMarkWhen(out io.Writer) *MarkWhen {
	return &MarkWhen{out}
}

// Generate generates Markwhen compatible fact export
func (mw MarkWhen) Generate(events biograph.Events) error {
	// fmt.Fprintln(mw.out, "dateFormat: y-M-d")
	for _, event := range events {
		if err := mw.printEvent(event); err != nil {
			return err
		}
	}
	return nil
}

func formatMarkWhenTime(t time.Time) string {
	if t.IsZero() {
		return "now"
	}
	return t.Format("2006-01-02")
}

func (mw MarkWhen) printEvent(e biograph.LifeEvent) error {
	if _, err := fmt.Fprintf(mw.out, "%s - %s: %s #%s\n", formatMarkWhenTime(e.GetFrom()), formatMarkWhenTime(e.GetTo()), e.GetName(), e.GetType()); err != nil {
		return err
	}
	return nil
}

package report

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/aquilax/biograph"
)

const dateFormat = "2006-01-02"

type Text struct {
	out io.Writer
}

// NewText creates new text reporter
func NewText(out io.Writer) *Text {
	return &Text{out}
}

// Generate generates chronological text report
func (r *Text) Generate(events biograph.Events) error {
	for _, event := range events {
		if err := r.printEvent(event); err != nil {
			return err
		}
	}
	return nil
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return "----------"
	}
	return t.Format(dateFormat)
}

func (r *Text) printEvent(e biograph.LifeEvent) error {
	_, err := fmt.Fprintf(r.out, "%s - %s %s %s (%s)\n", formatTime(e.GetFrom()), formatTime(e.GetTo()), getTypeSymbol(e.GetType()), e.GetName(), renderMeta(e.GetMeta()))
	return err
}

func getTypeSymbol(et biograph.EventType) string {
	switch et {
	case biograph.Home:
		return "🏠"
	case biograph.Education:
		return "🏫"
	case biograph.Work:
		return "👷"
	case biograph.Travel:
		return "✈️"
	case biograph.Item:
		return "📦"
	case biograph.Partner:
		return "❤️"
	case biograph.Roommate:
		return "😃"
	case biograph.Project:
		return "💡"
	case biograph.Document:
		return "🗎"
	case biograph.Process:
		return "⚙️"
	}
	return "?"
}

func renderMeta(m *biograph.MetaData) string {
	list := make([]string, len(*m))
	i := 0
	keys := m.Keys()
	sort.StringSlice.Sort(keys)
	for _, key := range keys {
		list[i] = fmt.Sprintf("%s=%s", key, m.Get(key))
		i++
	}
	return strings.Join(list, ", ")
}

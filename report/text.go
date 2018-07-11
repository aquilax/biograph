package report

import (
	"fmt"
	"io"
	"strings"

	"github.com/aquilax/biograph"
)

const dateFormat = "2006-01-02"

type Text struct {
	out io.WriteCloser
}

// NewText creates new text reporter
func NewText(out io.WriteCloser) *Text {
	return &Text{out}
}

// Generate generates chronological text report
func (r *Text) Generate(life *biograph.Life) error {
	for _, event := range life.Asc() {
		if err := r.printEvent(event); err != nil {
			return err
		}
	}
	return nil
}

func (r *Text) printEvent(e biograph.LifeEvent) error {
	_, err := fmt.Fprintf(r.out, "%s - %s %s %s (%s)\n", e.GetFrom().Format(dateFormat), e.GetTo().Format(dateFormat), getTypeSymbol(e.GetType()), e.GetName(), renderMeta(e.GetMeta()))
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
	}
	return "?"
}

func renderMeta(m *biograph.MetaData) string {
	list := make([]string, len(*m))
	i := 0
	for key, val := range *m {
		list[i] = fmt.Sprintf("%s=%s", key, val)
		i++
	}
	return strings.Join(list, ", ")
}

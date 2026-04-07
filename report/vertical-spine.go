package report

import (
	"fmt"
	"io"
	"strings"

	"github.com/aquilax/biograph"
)

const columnWidth = 38
const gutter = 2
const spineChar = "|"

type VerticalSpine struct {
	out io.Writer
}

func NewVerticalSpine(out io.Writer) *VerticalSpine {
	return &VerticalSpine{out}
}

func (r *VerticalSpine) Generate(events []biograph.LifeEvent) error {
	for i, event := range events {
		date := event.GetFrom().Format(dateFormat)
		name := event.GetName()
		lineText := fmt.Sprintf("%s %s", date, name)

		var line string
		if i%2 == 0 {
			// Left side event
			// Logic: [Text] [Underline] [Spine]
			padding := columnWidth - len(lineText) - gutter
			if padding < 0 {
				padding = 0
			}
			line = fmt.Sprintf("%s%s%s\n",
				strings.Repeat(" ", padding),
				lineText,
				strings.Repeat("_", gutter)+spineChar)
		} else {
			// Right side event
			// Logic: [Spine] [Underline] [Text]
			line = fmt.Sprintf("%s%s %s\n",
				strings.Repeat(" ", columnWidth)+spineChar,
				strings.Repeat("_", gutter),
				lineText)
		}

		if _, err := fmt.Fprint(r.out, line); err != nil {
			return err
		}

		// Add some vertical space (the spine) between events
		fmt.Fprintf(r.out, "%s%s\n", strings.Repeat(" ", columnWidth), spineChar)
	}
	return nil
}

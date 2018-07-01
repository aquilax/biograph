package biograph

import (
	"fmt"
	"io"
)

const dateFormat = "2006-01-02"

type RenderOptions map[string]string

type Renderer interface {
	render(*Life) error
}

type TextRenderer struct {
	out io.WriteCloser
}

func NewTextRenderer(out io.WriteCloser) *TextRenderer {
	return &TextRenderer{out}
}

func (tr *TextRenderer) Render(life *Life) error {
	for _, event := range life.Asc() {
		fmt.Fprintf(tr.out, "%s - %s  %s\n", event.getFrom().Format(dateFormat), event.getTo().Format(dateFormat), event.getName())
	}
	return nil
}

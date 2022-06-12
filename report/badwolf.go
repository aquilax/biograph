package report

import (
	"fmt"
	"io"

	"github.com/aquilax/biograph"
)

type BadWolf struct {
	out io.Writer
}

func NewBadWolf(out io.Writer) *BadWolf {
	return &BadWolf{out}
}

// Generate generates badwolf compatible fact export
func (bw BadWolf) Generate(events biograph.Events) error {
	for _, event := range events {
		if err := bw.printEvent(event); err != nil {
			return err
		}
	}
	return nil
}

func (bw BadWolf) printEvent(e biograph.LifeEvent) error {
	if !e.GetFrom().IsZero() {
		if _, err := fmt.Fprintf(bw.out, "/%s<%s> \"%s\"@[%s] /%s<%s>\n", "person", "me", "start", e.GetFrom(), e.GetType(), e.GetName()); err != nil {
			return err
		}
	}
	if !e.GetTo().IsZero() {
		if _, err := fmt.Fprintf(bw.out, "/%s<%s> \"%s\"@[%s] /%s<%s>\n", "person", "me", "stop", e.GetTo(), e.GetType(), e.GetName()); err != nil {
			return err
		}
	}

	meta := e.GetMeta()
	if meta != nil {
		for k, v := range *meta {
			if k != "name" {
				if _, err := fmt.Fprintf(bw.out, "/%s<%s> \"%s\"@[] /%s<%s>\n", e.GetType(), e.GetName(), "is", k, v); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

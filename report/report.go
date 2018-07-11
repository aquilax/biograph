package report

import (
	"github.com/aquilax/biograph"
)

type Report interface {
	Generate(l *biograph.Life) error
}

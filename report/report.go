package report

import (
	"github.com/aquilax/biograph"
)

type Report interface {
	Generate([]biograph.LifeEvent) error
}

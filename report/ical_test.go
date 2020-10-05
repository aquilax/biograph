package report

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/aquilax/biograph"
)

func TestNewICal(t *testing.T) {
	life := biograph.NewLifeArray(d("1900-01-01"), d("1990-01-01"))
	life.Add(
		biograph.NewHome("City", "Country", d("1900-01-01"), d("1990-01-01"), nil),
	)

	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	r := NewICal(bw)
	r.Generate(life.Items())
	bw.Flush()

	want := `BEGIN:VCALENDAR
PRODID:-//Biograph//NONSGML Calendar//EN
VERSION:2.0
BEGIN:VEVENT
CATEGORIES:home
DTSTAMP:19000101T000000Z
DTSTART:19900101T000000Z
SUMMARY:City
UID:uid@example.org
END:VEVENT
END:VCALENDAR
`
	got := b.String()
	if !reflect.DeepEqual(len(strings.Split(got, "\n")), len(strings.Split(want, "\n"))) {
		t.Errorf("NewLifeCalendar() = %v, want %v", got, want)
	}
}

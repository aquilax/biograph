package report

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/aquilax/biograph"
)

func d(date string) time.Time {
	time, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		panic(err)
	}
	return time
}

func TestNewLifeCalendar(t *testing.T) {

	life := biograph.NewLifeArray(d("1900-01-01"), d("1990-01-01"))
	life.Add(
		biograph.NewHome("City", "Country", d("1900-01-01"), d("1990-01-01"), nil),
	)

	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	r := NewLifeCalendar(bw)
	r.Generate(life.Items())
	bw.Flush()

	want := `{"options":null,"events":[{"type":1,"date":"1900-01-01","title":"City"}]}`
	got := b.String()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewLifeCalendar() = %v, want %v", got, want)
	}
}

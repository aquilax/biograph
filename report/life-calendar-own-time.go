package report

import (
	"fmt"
	"io"
	"time"

	"github.com/aquilax/biograph"
)

type LifeCalendarOwnTime struct {
	startDate time.Time
	endDate   time.Time
	output    io.Writer
}

func NewLifeCalendarOwnTime(startDate, endDate time.Time, o io.Writer) *LifeCalendarOwnTime {
	return &LifeCalendarOwnTime{startDate, endDate, o}
}

func (l *LifeCalendarOwnTime) Generate(events biograph.Events) error {
	fromYear := 0
	toYear := int(l.endDate.Sub(l.startDate).Hours()/(24*7*52)) + 1
	buckets := getWeekBucketsOwnTime(l.startDate, events)

	for year := fromYear; year < toYear; year++ {
		totalWeeks := 52
		fmt.Fprintf(l.output, "% 3d  ", year+1)
		for week := 0; week < totalWeeks; week++ {
			key := bucketKey(year, week)
			if count, ok := buckets[key]; ok {
				fmt.Fprintf(l.output, "%02d", count)
			} else {
				fmt.Fprint(l.output, "--")
			}
			if week < totalWeeks-1 {
				l.output.Write([]byte{' '})
			}
		}
		fmt.Fprintln(l.output, "")
	}
	return nil
}

func getWeekBucketsOwnTime(startDate time.Time, events biograph.Events) map[string]int {
	buckets := make(map[string]int)
	for _, event := range events {
		daysFromStart := int(event.GetFrom().Sub(startDate).Hours() / 24)
		weeksFromStart := int(daysFromStart / 7)
		year := weeksFromStart / 52
		week := weeksFromStart - (int(weeksFromStart/52) * 52)
		key := bucketKey(year, week)
		buckets[key] += 1
	}
	return buckets
}

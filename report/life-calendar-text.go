package report

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/aquilax/biograph"
)

type LifeCalendarText struct {
	startDate time.Time
	endDate   time.Time
	output    io.Writer
}

func NewLifeCalendarText(startDate, endDate time.Time, o io.Writer) *LifeCalendarText {
	return &LifeCalendarText{startDate, endDate, o}
}

func (l *LifeCalendarText) Generate(events biograph.Events) error {
	fromYear := l.startDate.Year()
	toYear := l.endDate.Year()
	buckets := getWeekBuckets(events)

	for year := fromYear; year <= toYear; year++ {
		totalWeeks := weeksInYear(year)
		fmt.Fprintf(l.output, "%04d  ", year)
		for week := 1; week <= totalWeeks; week++ {
			key := bucketKey(year, week)
			if count, ok := buckets[key]; ok {
				fmt.Fprintf(l.output, "%02d", count)
			} else {
				fmt.Fprintf(l.output, "%02d", 0)
			}
			if week < totalWeeks {
				l.output.Write([]byte{' '})
			}
			//fmt.Fprintf(l.output, "%02d ", week)
		}
		fmt.Fprintln(l.output, "")
	}
	return nil
}

func weeksInYear(year int) int {
	lastDay := time.Date(year, 12, 28, 1, 1, 1, 1, time.Now().Location())
	_, week := lastDay.ISOWeek()
	return week
}

func bucketKey(year, week int) string {
	return strconv.Itoa(year) + "-" + strconv.Itoa(week)
}

func getWeekBuckets(events biograph.Events) map[string]int {
	buckets := make(map[string]int)
	for _, event := range events {
		key := bucketKey(event.GetFrom().ISOWeek())
		buckets[key] += 1
	}
	return buckets
}

package report

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/aquilax/biograph"
)

// Branches renders a git-like ASCII branches/merges timeline for life events
// It shows events starting with '*' and continuing with '|', and marks ends with '┤'.
// The reporter optionally filters events to those overlapping [startDate, endDate].
type Branches struct {
	startDate time.Time
	endDate   time.Time
	out       io.Writer
}

func NewBranches(startDate, endDate time.Time, out io.Writer) *Branches {
	return &Branches{startDate: startDate, endDate: endDate, out: out}
}

// Generate prints a simple ASCII graph of branches/merges based on events' from/to dates.
func (b *Branches) Generate(events biograph.Events) error {
	// filter events by overlap with provided range
	filtered := make([]biograph.LifeEvent, 0)
	for _, e := range events {
		if overlaps(e.GetFrom(), e.GetTo(), b.startDate, b.endDate) {
			filtered = append(filtered, e)
		}
	}

	if len(filtered) == 0 {
		_, err := fmt.Fprintln(b.out, "No events in range")
		return err
	}

	// collect unique dates (starts and ends)
	starts := make(map[time.Time][]int)
	ends := make(map[time.Time][]int)
	datesMap := make(map[time.Time]struct{})

	for i, e := range filtered {
		from := toDateOnly(e.GetFrom())
		starts[from] = append(starts[from], i)
		datesMap[from] = struct{}{}
		if !e.GetTo().IsZero() {
			end := toDateOnly(e.GetTo())
			ends[end] = append(ends[end], i)
			datesMap[end] = struct{}{}
		}
	}

	// sort dates
	dates := make([]time.Time, 0, len(datesMap))
	for d := range datesMap {
		dates = append(dates, d)
	}
	sort.Slice(dates, func(i, j int) bool { return dates[i].Before(dates[j]) })

	// columns represent active event indices (into filtered)
	columns := make([]int, 0)

	for _, d := range dates {
		// add new events that start now (append to right)
		if ids, ok := starts[d]; ok {
			for _, id := range ids {
				columns = append(columns, id)
			}
		}

		// build line
		chars := make([]string, len(columns))
		// map event id -> column index for connector placement
		idxOf := make(map[int]int)
		for i, id := range columns {
			idxOf[id] = i
			evt := filtered[id]
			if toDateOnly(evt.GetFrom()).Equal(d) {
				chars[i] = "*"
			} else if !evt.GetTo().IsZero() && toDateOnly(evt.GetTo()).Equal(d) {
				chars[i] = "┤"
			} else {
				chars[i] = "|"
			} 
		}

		// add branch connectors (├):
		// for every event starting now, mark the previous column with '├' if available
		if ids, ok := starts[d]; ok {
			for _, sid := range ids {
				colIdx := idxOf[sid]
				if colIdx > 0 {
					// don't overwrite start or end markers
					if chars[colIdx-1] != "*" && chars[colIdx-1] != "┤" {
						chars[colIdx-1] = "├"
					}
				}
			}
		}

		// for every event ending now, mark previous column with merge connector '┤'
		if ids, ok := ends[d]; ok {
			for _, eid := range ids {
				colIdx := idxOf[eid]
				if colIdx > 0 {
					if chars[colIdx] == "┤" {
						if chars[colIdx-1] != "*" && chars[colIdx-1] != "┤" {
							chars[colIdx-1] = "┤"
						}
					}
				}
			}
		}

		line := fmt.Sprintf("%s  %s", d.Format("2006-01-02"), strings.Join(chars, " "))

		// append names of newly started events after the graph
		names := make([]string, 0)
		for _, id := range columns {
			evt := filtered[id]
			if toDateOnly(evt.GetFrom()).Equal(d) {
				names = append(names, fmt.Sprintf("%s (%s)", evt.GetName(), evt.GetType()))
			}
		}
		if len(names) > 0 {
			line = line + "  " + strings.Join(names, ", ")
		}

		if _, err := fmt.Fprintln(b.out, line); err != nil {
			return err
		}

		// remove events that end now
		if ids, ok := ends[d]; ok {
			if len(ids) == 0 {
				continue
			}
			// build set for quick lookup
			endSet := make(map[int]struct{})
			for _, id := range ids {
				endSet[id] = struct{}{}
			}
			newCols := make([]int, 0, len(columns))
			for _, cid := range columns {
				if _, ending := endSet[cid]; !ending {
					newCols = append(newCols, cid)
				}
			}
			columns = newCols
		}
	}

	return nil
}

func overlaps(aFrom, aTo, bFrom, bTo time.Time) bool {
	// treat zero aTo as ongoing
	if aTo.IsZero() {
		// overlaps if aFrom <= bTo (or bTo zero) and aFrom >= bFrom? use simple check: aFrom <= bTo (if bTo non-zero) OR bTo zero
		if bTo.IsZero() {
			return !aFrom.After(bFrom)
		}
		return !aFrom.After(bTo) && !aTo.Before(bFrom)
	}

	// general interval overlap
	if bTo.IsZero() {
		return !aTo.Before(bFrom) && !aFrom.After(bTo)
	}
	return !(aTo.Before(bFrom) || aFrom.After(bTo))
}

// toDateOnly normalizes time to year-month-day with zeroed clock in UTC
func toDateOnly(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

package report

import (
	"io"
	"math"
	"sort"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/aquilax/biograph"
)

type canvasEvent struct {
	lane int
	biograph.LifeEvent
}

type eventGroup struct {
	eventType biograph.EventType
	items     []canvasEvent
	maxSize   int
}

type SVG struct {
	out        io.WriteCloser
	scale      float64
	laneHeight int
	minTime    time.Time
	maxTime    time.Time
	groups     map[biograph.EventType]*eventGroup
	canvas     *svg.SVG
}

func NewSVG(out io.WriteCloser) *SVG {
	return &SVG{
		scale:      1,
		laneHeight: 16,
		groups:     make(map[biograph.EventType]*eventGroup),
		canvas:     svg.New(out),
	}
}

func (s *SVG) Generate(le biograph.Events) error {
	s.init(le)
	s.canvas.Startpercent(100, 100)
	s.drawItems()
	s.canvas.Style("text/css", `
.event-label {font-family: Arial; font-weight: 400; font-size: 12px; alignment-baseline:middle; text-anchor:middle }
.lane {fill: #fff; stroke: #000; stroke-width: 1}
.event {fill: #f00; stroke: #000; fill-opacity:0.4; stroke-width: 1}
	`)
	s.canvas.End()
	return nil
}

func (s *SVG) init(le biograph.Events) {
	// get min time
	le.Sort(biograph.AscFrom)
	s.minTime = le[0].GetFrom()
	// get max time
	le.Sort(biograph.DescTo)
	s.maxTime = le[0].GetTo()

	// group events by type
	s.groupItems(le)
}

func (s *SVG) groupItems(le biograph.Events) {
	var et biograph.EventType
	for _, event := range le {
		et = event.GetType()
		if _, ok := s.groups[et]; !ok {
			s.groups[et] = &eventGroup{
				eventType: event.GetType(),
				items:     make([]canvasEvent, 0),
				maxSize:   0,
			}
		}
		s.groups[et].items = append(s.groups[et].items, canvasEvent{-1, event})
	}
}

func arrange(ce []canvasEvent) int {
	result := 0
	// sort first
	sort.Slice(ce, func(i, j int) bool { return ce[i].GetFrom().Before(ce[j].GetFrom()) })
Loop:
	for n := range ce {
		if n == 0 {
			ce[n].lane = 0
			continue
		}
		maxLane := 0
		for i := 0; i < n; i++ {
			maxLane = ce[i].lane
			if ce[n].GetFrom().Sub(ce[i].GetTo()) >= 0 {
				ce[n].lane = ce[i].lane
				break Loop
			}
		}
		ce[n].lane = maxLane + 1
		if ce[n].lane > result {
			result = ce[n].lane
		}
	}
	return result
}

func (s *SVG) drawItems() {
	offset := 0
	for _, group := range s.groups {
		s.canvas.Translate(0, offset)
		maxLane := arrange(group.items)
		offset = s.drawPool(maxLane+1, group)
		for _, i := range group.items {
			x := diffScale(s.minTime, i.GetFrom(), s.scale)
			y := i.lane * s.laneHeight
			h := s.laneHeight
			w := diffScale(i.GetFrom(), i.GetTo(), s.scale)
			s.canvas.Translate(x, y)
			s.canvas.Rect(0, 0, w, h, "class=\"event\"")
			s.canvas.Text(int(w/2), int(h/2), i.GetName(), "class=\"event-label\"")
			s.canvas.Gend()
		}
		s.canvas.Gend()
	}
}

func diffScale(from, to time.Time, scale float64) int {
	return int(math.Ceil(to.Sub(from).Hours() / 24.0 * scale))
}

func (s *SVG) drawPool(maxLane int, g *eventGroup) int {
	x := 0
	y := 0
	w := diffScale(s.minTime, s.maxTime, s.scale)
	h := maxLane * s.laneHeight
	s.canvas.Translate(x, y)
	s.canvas.Rect(0, 0, w, h, "class=\"lane\"")
	s.canvas.Gend()
	return h
}

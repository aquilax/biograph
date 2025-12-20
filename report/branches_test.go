package report

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/aquilax/biograph"
)

func mustParse(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}

func TestBranches_SimpleBranching(t *testing.T) {
	// A runs all year 2020
	// B starts in May 2020 overlapping A
	// C starts in Aug 2020 overlapping B
	a := biograph.NewWork("CompanyA", "Engineer", mustParse("2020-01-01"), mustParse("2020-12-31"), nil)
	b := biograph.NewProject("ProjectB", mustParse("2020-05-01"), mustParse("2021-05-01"), nil)
	c := biograph.NewTravel("Paris", "France", mustParse("2020-08-01"), mustParse("2020-10-01"), nil)

	events := biograph.Events{a, b, c}

	var out bytes.Buffer
	rep := NewBranches(mustParse("2020-01-01"), mustParse("2021-06-01"), &out)
	if err := rep.Generate(events); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	s := out.String()
	// Expect these lines to exist
	expects := []string{
		"2020-01-01", // A starts
		"CompanyA (work)", // name and type
		"2020-05-01", // B starts
		"ProjectB (project)", // B start
		"├", // branch connector when branching
		"2020-08-01", // C starts
		"Paris, France (travel)", // C start (travel event name uses place)
		"2020-10-01", // C ends
		"┤", // merge connector on end
	}

	for _, e := range expects {
		if !strings.Contains(s, e) {
			t.Fatalf("expected output to contain %q, got:\n%s", e, s)
		}
	}
}

func TestBranches_OutOfRangeFiltered(t *testing.T) {
	a := biograph.NewWork("CompanyA", "Engineer", mustParse("2019-01-01"), mustParse("2019-12-31"), nil)
	b := biograph.NewProject("ProjectB", mustParse("2020-05-01"), mustParse("2020-07-01"), nil)
	events := biograph.Events{a, b}

	var out bytes.Buffer
	rep := NewBranches(mustParse("2020-01-01"), mustParse("2020-12-31"), &out)
	if err := rep.Generate(events); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	s := out.String()
	if !strings.Contains(s, "2020-05-01") || !strings.Contains(s, "ProjectB") {
		t.Fatalf("expected filtered report for ProjectB but got:\n%s", s)
	}
	if strings.Contains(s, "CompanyA") {
		t.Fatalf("did not expect CompanyA to appear in report, got:\n%s", s)
	}
}

package biograph_test

import (
	"fmt"
	"time"

	bio "github.com/aquilax/biograph"
)

func mustDate(date string) time.Time {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return time
}

func ExampleBiograph() {
	l := bio.NewLife(mustDate("1980-01-01"), time.Now())
	l.Add(bio.NewHome("Example Avenue, My City", "My Country", mustDate("2001-01-01"), mustDate("2010-02-02"), nil))
	l.Add(bio.NewEducation("My School", "Bachelor", mustDate("2010-01-01"), mustDate("2015-02-02"), nil))
	l.Add(bio.NewWork("My Employer", "Employee", mustDate("2010-01-01"), mustDate("2015-02-02"), nil))
	fmt.Println(l.Count())
	// Output: 3
}

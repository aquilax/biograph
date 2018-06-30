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
	l.Add(bio.NewEvent())
	fmt.Println(l.Count())
	// Output: 1
}

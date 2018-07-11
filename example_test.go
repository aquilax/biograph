package biograph_test

import (
	"os"
	"time"

	bio "github.com/aquilax/biograph"
	"github.com/aquilax/biograph/report"
)

func mustDate(date string) time.Time {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return time
}

func ExampleLife() {
	l := bio.NewLife(mustDate("1980-01-01"), time.Now())
	l.Add(
		bio.NewHome("Example Avenue, My City", "My Country", mustDate("2001-01-01"), mustDate("2010-02-02"), nil),
		bio.NewEducation("My School", "Bachelor", mustDate("2002-01-01"), mustDate("2015-02-02"), nil),
		bio.NewWork("My Employer", "Employee", mustDate("2003-01-01"), mustDate("2015-02-02"), nil),
		bio.NewTravel("My Place", "My Country", mustDate("2004-01-01"), mustDate("2013-02-02"), nil),
		bio.NewItem("test/category", mustDate("2008-01-01"), mustDate("2015-02-02"), &bio.MetaData{"brand": "Brand"}),
	)

	tr := report.NewText(os.Stdout)
	tr.Generate(l)
	// Output:
	// 2001-01-01 - 2010-02-02 🏠 Example Avenue, My City
	// 2002-01-01 - 2015-02-02 🏫 My School
	// 2003-01-01 - 2015-02-02 👷 My Employer
	// 2004-01-01 - 2013-02-02 ✈ My Place, My Country
	// 2008-01-01 - 2015-02-02 📦 test/category
}

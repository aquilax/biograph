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
		bio.NewPartner("My Partner", mustDate("2012-01-01"), mustDate("2015-03-02"), nil),
	)

	tr := report.NewText(os.Stdout)
	tr.Generate(l)
	// Output:
	// 2001-01-01 - 2010-02-02 🏠 Example Avenue, My City (address=Example Avenue, My City, country=My Country)
	// 2002-01-01 - 2015-02-02 🏫 My School (degree=Bachelor, school=My School)
	// 2003-01-01 - 2015-02-02 👷 My Employer (employer=My Employer, position=Employee)
	// 2004-01-01 - 2013-02-02 ✈️ My Place, My Country (country=My Country, place=My Place)
	// 2008-01-01 - 2015-02-02 📦 test/category (brand=Brand, category=test/category)
	// 2012-01-01 - 2015-03-02 ❤️ My Partner (name=My Partner)
}

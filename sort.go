package biograph

type CompGenerator func(Events) Comparator
type Comparator func(int, int) bool

func Asc(le Events) Comparator {
	return func(i, j int) bool {
		return le[i].GetFrom().Before(le[j].GetFrom())
	}
}

func Desc(le Events) Comparator {
	return func(i, j int) bool {
		return le[j].GetFrom().Before(le[i].GetFrom())
	}
}

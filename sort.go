package biograph

type CompGenerator func(Events) Comparator
type Comparator func(int, int) bool

func AscFrom(le Events) Comparator {
	return func(i, j int) bool {
		return le[i].GetFrom().Before(le[j].GetFrom())
	}
}

func DescFrom(le Events) Comparator {
	return func(i, j int) bool {
		return le[j].GetFrom().Before(le[i].GetFrom())
	}
}

func AscTo(le Events) Comparator {
	return func(i, j int) bool {
		return le[i].GetTo().Before(le[j].GetTo())
	}
}

func DescTo(le Events) Comparator {
	return func(i, j int) bool {
		return le[j].GetTo().Before(le[i].GetTo())
	}
}

func AscName(le Events) Comparator {
	return func(i, j int) bool {
		return le[i].GetName() < le[j].GetName()
	}
}

func DescName(le Events) Comparator {
	return func(i, j int) bool {
		return le[i].GetName() > le[j].GetName()
	}
}

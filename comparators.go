package cardobject

type comparators int

const (
	EQUAL comparators = iota
	GREATER
	SMALLER	
)

type Comparator interface {
	Comparators() comparators
}

func (s comparators) Comparators() comparators {
	return s
}
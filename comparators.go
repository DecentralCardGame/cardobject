package cardobject

const (
	EQUAL comparators = iota
	GREATER
	SMALLER	
)

type Comparator interface {
	Comparators() comparators
}


type comparators int

func (s comparators) Comparators() comparators {
	return s
}
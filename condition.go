package cardobject

type Condition struct {
	property string
	value int
	comparator Comparator
}

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
package cardobject

type arithOperator int

const (
	ADD arithOperator = iota
	SUB
)

type ArithOperator interface {
	ArithOperator() arithOperator
}

func (a arithOperator) ArithOperator() arithOperator {
	return a
}
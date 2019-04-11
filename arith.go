package cardobject

type arithOperator int

const (
	Set arithOperator = iota
	ADD
	SUB
)

type ArithOperator interface {
	ArithOperator() arithOperator
}

func (a arithOperator) ArithOperator() arithOperator {
	return a
}
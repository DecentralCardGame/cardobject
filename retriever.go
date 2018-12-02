package cardobject

type IntInserter interface {
	IsIntRetriever() bool
}
type StringInserter interface {
	IsStringRetriever() bool
}


func NewIntConst(i int) IntInserter {
	return &intConst{i}
}

func NewStringConst(s string) StringInserter {
	return &stringConst{s}
}


type intRetriever struct {
	prop IntPropertyId
}

type stringRetriever struct {
	prop StringPropertyId
}

type intConst struct {
	i int
}

type stringConst struct {
	s string
}


func (ir *cardIntPropertyId) IsIntRetriever() bool {
	return true
}

func (ir *playerIntPropertyId) IsIntRetriever() bool {
	return true
}

func (sr cardStringPropertyId) IsStringRetriever() bool {
	return true
}

func (sr playerStringPropertyId) IsStringRetriever() bool {
	return true
}

func (ic *intConst) IsIntRetriever() bool {
	return true
}

func (sc *stringConst) IsStringRetriever() bool {
	return true
}
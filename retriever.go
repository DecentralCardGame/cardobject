package cardobject

type IntInserter interface {
	RetrieveInt() int
}
type StringInserter interface {
	RetrieveString() string
}

func NewIntRetriever(cip CardIntProperty) IntInserter {
	return &intRetriever{cip}
}
func NewStringRetriever(csp CardStringProperty) StringInserter {
	return &stringRetriever{csp}
}

func NewIntConst(i int) IntInserter {
	return &intConst{i}
}

func NewStringConst(s string) StringInserter {
	return &stringConst{s}
}


type intRetriever struct {
	prop CardIntProperty
}

type stringRetriever struct {
	prop CardStringProperty
}

type intConst struct {
	i int
}

type stringConst struct {
	s string
}


func (ir *intRetriever) RetrieveInt() int {
	return ir.prop.ExtractIntProp()
}

func (sr *stringRetriever) RetrieveString() string {
	return sr.prop.ExtractStringProp()
}

func (ic *intConst) RetrieveInt() int {
	return ic.i
}

func (sc *stringConst) RetrieveString() string {
	return sc.s
}
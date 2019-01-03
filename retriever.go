package cardobject

type IntInserter interface {
	IsIntRetriever() bool
}
type StringInserter interface {
	IsStringRetriever() bool
}

type IntInserterPropId interface {
	IntInserter
	GetIntPropertyId() IntPropertyId
}

type StringInserterPropId interface {
	StringInserter
	GetStringPropertyId() StringPropertyId
}



func NewIntConst(i int) *IntInserterConst {
	return &IntInserterConst{i}
}

func NewStringConst(s string) *StringInserterConst {
	return &StringInserterConst{s}
}


type IntInserterConst struct {
	i int
}

type StringInserterConst struct {
	s string
}


func (cipi CardIntPropertyId) IsIntRetriever() bool {
	return true
}

func (pipi PlayerIntPropertyId) IsIntRetriever() bool {
	return true
}

func (cspi CardStringPropertyId) IsStringRetriever() bool {
	return true
}

func (pspi PlayerStringPropertyId) IsStringRetriever() bool {
	return true
}

func (ic *IntInserterConst) IsIntRetriever() bool {
	return true
}

func (sc *StringInserterConst) IsStringRetriever() bool {
	return true
}

func (cipi CardIntPropertyId) GetIntPropertyId() CardIntPropertyId {
	return cipi
}

func (cspi CardStringPropertyId) GetStringPropertyId() CardStringPropertyId {
	return cspi
}

func (pipi PlayerIntPropertyId) GetIntPropertyId() PlayerIntPropertyId {
	return pipi
}

func (pspi PlayerStringPropertyId) GetStringPropertyId() PlayerStringPropertyId {
	return pspi
}

func (ic *IntInserterConst) GetIntVal() int {
	return ic.i
}

func (sc *StringInserterConst) GetStringVal() string {
	return sc.s
}
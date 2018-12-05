package cardobject

type IntInserter interface {
	IsIntRetriever() bool
}
type StringInserter interface {
	IsStringRetriever() bool
}

type IntInserterConst interface {
	IntInserter
	GetIntVal() int
}

type StringInserterConst interface {
	StringInserter
	GetStringVal() string
}

type IntInserterPropId interface {
	IntInserter
	GetIntPropertyId() IntPropertyId
}

type StringInserterPropId interface {
	StringInserter
	GetStringPropertyId() StringPropertyId
}



func NewIntConst(i int) IntInserterConst {
	return &intConst{i}
}

func NewStringConst(s string) StringInserterConst {
	return &stringConst{s}
}


type intConst struct {
	i int
}

type stringConst struct {
	s string
}


func (cipi cardIntPropertyId) IsIntRetriever() bool {
	return true
}

func (pipi playerIntPropertyId) IsIntRetriever() bool {
	return true
}

func (cspi cardStringPropertyId) IsStringRetriever() bool {
	return true
}

func (pspi playerStringPropertyId) IsStringRetriever() bool {
	return true
}

func (ic *intConst) IsIntRetriever() bool {
	return true
}

func (sc *stringConst) IsStringRetriever() bool {
	return true
}

func (cipi cardIntPropertyId) GetIntPropertyId() cardIntPropertyId {
	return cipi
}

func (cspi cardStringPropertyId) GetStringPropertyId() cardStringPropertyId {
	return cspi
}

func (pipi playerIntPropertyId) GetIntPropertyId() playerIntPropertyId {
	return pipi
}

func (pspi playerStringPropertyId) GetStringPropertyId() playerStringPropertyId {
	return pspi
}

func (ic *intConst) GetIntVal() int {
	return ic.i
}

func (sc *stringConst) GetStringVal() string {
	return sc.s
}
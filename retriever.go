package cardobject

type Inserter interface {
	IsRetriever() bool
}

type IntInserter interface {
	Inserter
	IsIntRetriever() bool
}
type StringInserter interface {
	Inserter
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
	Val int
}

type stringConst struct {
	Val string
}

func (cipi cardIntPropertyId) IsRetriever() bool {
	return true
}

func (pipi playerIntPropertyId) IsRetriever() bool {
	return true
}

func (cspi cardStringPropertyId) IsRetriever() bool {
	return true
}

func (pspi playerStringPropertyId) IsRetriever() bool {
	return true
}

func (ic *intConst) IsRetriever() bool {
	return true
}

func (sc *stringConst) IsRetriever() bool {
	return true
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
	return ic.Val
}

func (sc *stringConst) GetStringVal() string {
	return sc.Val
}
package cardobject

type Condition interface {
	GetComparator() Comparator
}

type IntCondition interface {
	Condition
	GetCompVal() int
	GetIntPropertyId() IntPropertyId
}

type StringCondition interface {
	Condition
	GetCompVal() string
	GetStringPropertyId() StringPropertyId
}

type CardCondition interface {
	Condition
	GetCardPropertyId() CardPropertyId
}

type PlayerCondition interface {
	Condition
	GetPlayerPropertyId() PlayerPropertyId
}

func NewCardIntCondition(c Comparator, v int, p CardIntPropertyId) *CardIntCondition {
	return &CardIntCondition{&intCondition{&condition{c}, v}, p}
}

func NewCardStringCondition(v string, p CardStringPropertyId) *CardStringCondition {
	return &CardStringCondition{&stringCondition{&condition{EQUAL}, v}, p}
}

func NewPlayerIntCondition(c Comparator, v int, p PlayerIntPropertyId) *PlayerIntCondition {
	return &PlayerIntCondition{&intCondition{&condition{c}, v}, p}
}

func NewPlayerStringCondition(v string, p PlayerStringPropertyId) *PlayerStringCondition {
	return &PlayerStringCondition{&stringCondition{&condition{EQUAL}, v}, p}
}

type condition struct {
	Comparator Comparator
}

type intCondition struct{
	*condition
	Value int
}

type stringCondition struct{
	*condition
	Value string
}

type CardIntCondition struct {
	*intCondition
	Prop CardIntPropertyId
}

type CardStringCondition struct {
	*stringCondition
	Prop CardStringPropertyId
}

type PlayerIntCondition struct {
	*intCondition
	Prop PlayerIntPropertyId
}

type PlayerStringCondition struct {
	*stringCondition
	Prop PlayerStringPropertyId
}


func (c *condition) GetComparator() Comparator {
	return c.Comparator
}

func (ic *intCondition) GetCompVal() int {
	return ic.Value
}

func (sc *stringCondition) GetCompVal() string {
	return sc.Value
}

func (cic *CardIntCondition) GetCardPropertyId() CardPropertyId {
	return cic.Prop
}

func (cic *CardIntCondition) GetIntPropertyId() IntPropertyId {
	return cic.Prop
}

func (pic *PlayerIntCondition) GetPlayerPropertyId() PlayerPropertyId {
	return pic.Prop
}

func (pic *PlayerIntCondition) GetIntPropertyId() IntPropertyId {
	return pic.Prop
}

func (csc *CardStringCondition) GetCardPropertyId() CardPropertyId {
	return csc.Prop
}

func (csc *CardStringCondition) GetStringPropertyId() StringPropertyId {
	return csc.Prop
}

func (psc *PlayerStringCondition) GetPlayerPropertyId() PlayerPropertyId {
	return psc.Prop
}

func (psc *PlayerStringCondition) GetStringPropertyId() StringPropertyId {
	return psc.Prop
}
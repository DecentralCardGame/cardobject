package cardobject

type Condition interface {
	GetComparator() Comparator
}

type CardCondition interface {
	Condition
	GetCardPropertyId() CardPropertyId
}

type PlayerCondition interface {
	Condition
	GetPlayerPropertyId() PlayerPropertyId
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

func NewCardIntCondition(c Comparator, v int, p CardIntPropertyId) IntCondition {
	return &cardIntCondition{&intCondition{&condition{c}, v}, p}
}

func NewCardStringCondition(v string, p CardStringPropertyId) StringCondition {
	return &cardStringCondition{&stringCondition{&condition{EQUAL}, v}, p}
}

func NewPlayerIntCondition(c Comparator, v int, p PlayerIntPropertyId) IntCondition {
	return &playerIntCondition{&intCondition{&condition{c}, v}, p}
}

func NewPlayerStringCondition(v string, p PlayerStringPropertyId) StringCondition {
	return &playerStringCondition{&stringCondition{&condition{EQUAL}, v}, p}
}

type condition struct {
	comparator Comparator
}

type intCondition struct{
	*condition
	value int
}

type stringCondition struct{
	*condition
	value string
}

type cardIntCondition struct {
	*intCondition
	prop CardIntPropertyId
}

type cardStringCondition struct {
	*stringCondition
	prop CardStringPropertyId
}

type playerIntCondition struct {
	*intCondition
	prop PlayerIntPropertyId
}

type playerStringCondition struct {
	*stringCondition
	prop PlayerStringPropertyId
}


func (c *condition) GetComparator() Comparator {
	return c.comparator
}

func (ic *intCondition) GetCompVal() int {
	return ic.value
}

func (sc *stringCondition) GetCompVal() string {
	return sc.value
}

func (cic *cardIntCondition) GetCardPropertyId() CardPropertyId {
	return cic.prop
}

func (cic *cardIntCondition) GetIntPropertyId() IntPropertyId {
	return cic.prop
}

func (pic *playerIntCondition) GetPlayerPropertyId() PlayerPropertyId {
	return pic.prop
}

func (pic *playerIntCondition) GetIntPropertyId() IntPropertyId {
	return pic.prop
}

func (csc *cardStringCondition) GetCardPropertyId() CardPropertyId {
	return csc.prop
}

func (csc *cardStringCondition) GetStringPropertyId() StringPropertyId {
	return csc.prop
}

func (psc *playerStringCondition) GetPlayerPropertyId() PlayerPropertyId {
	return psc.prop
}

func (psc *playerStringCondition) GetStringPropertyId() StringPropertyId {
	return psc.prop
}
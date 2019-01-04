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

type CardIntCondition interface {
	GetCardPropertyId() CardPropertyId
	IntCondition
}

type CardStringCondition interface {
	GetCardPropertyId() CardPropertyId
	StringCondition
}

type PlayerIntCondition interface {
	GetPlayerPropertyId() PlayerPropertyId
	IntCondition
}

type PlayerStringCondition interface {
	GetPlayerPropertyId() PlayerPropertyId
	StringCondition
}

func NewCardIntCondition(c Comparator, v int, p CardIntPropertyId) CardIntCondition {
	return &cardIntCondition{&intCondition{&condition{c}, v}, p}
}

func NewCardStringCondition(v string, p CardStringPropertyId) CardStringCondition {
	return &cardStringCondition{&stringCondition{&condition{EQUAL}, v}, p}
}

func NewPlayerIntCondition(c Comparator, v int, p PlayerIntPropertyId) PlayerIntCondition {
	return &playerIntCondition{&intCondition{&condition{c}, v}, p}
}

func NewPlayerStringCondition(v string, p PlayerStringPropertyId) PlayerStringCondition {
	return &playerStringCondition{&stringCondition{&condition{EQUAL}, v}, p}
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

type cardIntCondition struct {
	*intCondition
	Prop CardIntPropertyId
}

type cardStringCondition struct {
	*stringCondition
	Prop CardStringPropertyId
}

type playerIntCondition struct {
	*intCondition
	Prop PlayerIntPropertyId
}

type playerStringCondition struct {
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

func (cic *cardIntCondition) GetCardPropertyId() CardPropertyId {
	return cic.Prop
}

func (cic *cardIntCondition) GetIntPropertyId() IntPropertyId {
	return cic.Prop
}

func (pic *playerIntCondition) GetPlayerPropertyId() PlayerPropertyId {
	return pic.Prop
}

func (pic *playerIntCondition) GetIntPropertyId() IntPropertyId {
	return pic.Prop
}

func (csc *cardStringCondition) GetCardPropertyId() CardPropertyId {
	return csc.Prop
}

func (csc *cardStringCondition) GetStringPropertyId() StringPropertyId {
	return csc.Prop
}

func (psc *playerStringCondition) GetPlayerPropertyId() PlayerPropertyId {
	return psc.Prop
}

func (psc *playerStringCondition) GetStringPropertyId() StringPropertyId {
	return psc.Prop
}
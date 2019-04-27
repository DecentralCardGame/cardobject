package cardobject

const (
	ATTACK cardIntPropertyId = iota
	COSTSUM
	HEALTH
	SPEEDMODIFIER

	NAME cardStringPropertyId = iota
	TAG 
	TEXT
	UNIQUENAME

	ENTITIESONBOARD playerIntPropertyId = iota
	DECKSIZE
	DUSTPILESIZE
	FIELDSONBOARD
	HANDSIZE
)

type Attack interface {
	CardIntProperty
	GetAttackVal() int
}

type attack struct {
	*cardIntProperty
}

func NewAttack(i int) Attack {
	return attack{&cardIntProperty{&intProperty{i}, ATTACK}}
}

func (a attack) GetAttackVal() int {
	return a.GetIntVal()
}


type CostSum interface {
	CardIntProperty
	GetCostSumVal() int
}

type costSum struct {
	*cardIntProperty
}

func NewCostSum(i int) CostSum {
	return costSum{&cardIntProperty{&intProperty{i}, COSTSUM}}
}

func (cs costSum) GetCostSumVal() int {
	return cs.GetIntVal()
}


type Health interface {
	CardIntProperty
	GetHealthVal() int
}

type health struct {
	*cardIntProperty
}

func NewHealth(i int) Health {
	return health{&cardIntProperty{&intProperty{i}, HEALTH}}
}

func (h health) GetHealthVal() int {
	return h.GetIntVal()
}


type Name interface {
	CardStringProperty
	GetNameVal() string
}

type name struct {
	*cardStringProperty
}

func NewName(s string) Name {
	return name{&cardStringProperty{&stringProperty{s}, NAME}}
}

func (n name) GetNameVal() string {
	return n.GetStringVal()
}


type SpeedModifier interface {
	CardIntProperty
	GetSpeedModifierVal() int
}

type speedModifier struct {
	*cardIntProperty
}

func NewSpeedModifier(i int) SpeedModifier {
	return speedModifier{&cardIntProperty{&intProperty{i}, SPEEDMODIFIER}}
}

func (sm speedModifier) GetSpeedModifierVal() int {
	return sm.GetIntVal()
}


type Tag interface {
	CardStringProperty
	GetTagVal() string
}

type tag struct {
	*cardStringProperty
}

func NewTag(s string) Tag {
	return tag{&cardStringProperty{&stringProperty{s}, TAG}}
}

func (t tag) GetTagVal() string {
	return t.GetStringVal()
}


type Text interface {
	CardStringProperty
	GetTextVal() string
}

type text struct {
	*cardStringProperty
}

func NewText(s string) Text {
	return text{&cardStringProperty{&stringProperty{s}, TEXT}}
}

func (t text) GetTextVal() string {
	return t.GetStringVal()
}


type UniqueName interface {
	CardStringProperty
	GetUniqueNameVal() string
}

type uniqueName struct {
	*name
}

func NewUniqueName(n string, d string) UniqueName {
	return uniqueName{&name{&cardStringProperty{&stringProperty{n + ", the " + d}, UNIQUENAME}}}
}

func (un uniqueName) GetUniqueNameVal() string {
	return un.GetStringVal()
}


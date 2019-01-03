package cardobject

const (
	ATTACK CardIntPropertyId = iota
	COSTSUM
	HEALTH
	SPEEDMODIFIER

	TAG CardStringPropertyId = iota
	TEXT

	ENTITIESONBOARD PlayerIntPropertyId = iota
	DECKSIZE
	DUSTPILESIZE
	FIELDSONBOARD
	HANDSIZE
)

type Attack struct {
	*CardIntProperty
}

func NewAttack(i int) *Attack {
	return &Attack{&CardIntProperty{&intProperty{i}, ATTACK}}
}

type CostSum struct {
	*CardIntProperty
}

func NewCostSum(i int) *CostSum {
	return &CostSum{&CardIntProperty{&intProperty{i}, COSTSUM}}
}

type Health struct {
	*CardIntProperty
}

func NewHealth(i int) *Health {
	return &Health{&CardIntProperty{&intProperty{i}, HEALTH}}
}

type Speedmodifier struct {
	*CardIntProperty
}

func NewSpeedmodifier(i int) *Speedmodifier {
	return &Speedmodifier{&CardIntProperty{&intProperty{i}, SPEEDMODIFIER}}
}

type Tag struct {
	*CardStringProperty
}

func NewTag(s string) *Tag {
	return &Tag{&CardStringProperty{&stringProperty{s}, TAG}}
}

type Text struct {
	*CardStringProperty
}

func NewText(s string) *Text {
	return &Text{&CardStringProperty{&stringProperty{s}, TEXT}}
}
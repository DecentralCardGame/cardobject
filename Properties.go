package cardobject

const (
	ATTACK cardIntPropertyId = iota
	COSTSUM
	HEALTH
	SPEEDMODIFIER

	NAME cardStringPropertyId = iota
	TAG 
	TEXT
	
	ENTITIESONBOARD playerIntPropertyId = iota
	DECKSIZE
	DUSTPILESIZE
	FIELDSONBOARD
	HANDSIZE
)

type Attack struct {
	*cardIntProperty
}

func NewAttack(i int) Attack {
	return Attack{&cardIntProperty{&intProperty{i}, ATTACK}}
}

type CostSum struct {
	*cardIntProperty
}

func NewCostSum(i int) CostSum {
	return CostSum{&cardIntProperty{&intProperty{i}, COSTSUM}}
}

type Health struct {
	*cardIntProperty
}

func NewHealth(i int) Health {
	return Health{&cardIntProperty{&intProperty{i}, HEALTH}}
}

type SpeedModifier struct {
	*cardIntProperty
}

func NewSpeedModifier(i int) SpeedModifier {
	return SpeedModifier{&cardIntProperty{&intProperty{i}, SPEEDMODIFIER}}
}

type Tag struct {
	*cardStringProperty
}

func NewTag(s string) Tag {
	return Tag{&cardStringProperty{&stringProperty{s}, TAG}}
}

type Text struct {
	*cardStringProperty
}

func NewText(s string) Text {
	return Text{&cardStringProperty{&stringProperty{s}, TEXT}}
}

type Name struct {
	*cardStringProperty
}

func NewName(s string) Name {
	return Name{&cardStringProperty{&stringProperty{s}, NAME}}
}
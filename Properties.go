package cardobject

const (
	ATTACK cardIntPropertyId = iota
	COSTSUM
	HEALTH
	SPEEDMODIFIER

	TAG cardStringPropertyId = iota
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

func NewAttack(i int) *Attack {
	return &Attack{&cardIntProperty{&intProperty{i}, ATTACK}}
}

type CostSum struct {
	*cardIntProperty
}

func NewCostSum(i int) *CostSum {
	return &CostSum{&cardIntProperty{&intProperty{i}, COSTSUM}}
}

type Health struct {
	*cardIntProperty
}

func NewHealth(i int) *Health {
	return &Health{&cardIntProperty{&intProperty{i}, HEALTH}}
}

type Speedmodifier struct {
	*cardIntProperty
}

func NewSpeedmodifier(i int) *Speedmodifier {
	return &Speedmodifier{&cardIntProperty{&intProperty{i}, SPEEDMODIFIER}}
}

type Tag struct {
	*cardStringProperty
}

func NewTag(s string) *Tag {
	return &Tag{&cardStringProperty{&stringProperty{s}, TAG}}
}

type Text struct {
	*cardStringProperty
}

func NewText(s string) *Text {
	return &Text{&cardStringProperty{&stringProperty{s}, TEXT}}
}


type EntitiesOnBoard struct {
	*playerIntProperty
}

func NewEntitiesOnBoard(i int) *EntitiesOnBoard {
	return &EntitiesOnBoard{&playerIntProperty{&intProperty{i}, ENTITIESONBOARD}}
}

type DeckSize struct {
	*playerIntProperty
}

func NewDeckSize(i int) *DeckSize {
	return &DeckSize{&playerIntProperty{&intProperty{i}, DECKSIZE}}
}

type DustpileSize struct {
	*playerIntProperty
}

func NewDustpileSize(i int) *DustpileSize {
	return &DustpileSize{&playerIntProperty{&intProperty{i}, DUSTPILESIZE}}
}

type FieldsOnBoard struct {
	*playerIntProperty
}

func NewFieldsOnBoard(i int) *FieldsOnBoard {
	return &FieldsOnBoard{&playerIntProperty{&intProperty{i}, FIELDSONBOARD}}
}

type HandSize struct {
	*playerIntProperty
}

func NewHandSize(i int) *HandSize {
	return &HandSize{&playerIntProperty{&intProperty{i}, HANDSIZE}}
}
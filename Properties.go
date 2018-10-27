package cardobject

type Attack struct {
	*cardIntProperty
}

func NewAttack(i int) *Attack {
	return &Attack{&cardIntProperty{&cardProperty{}, i}}
}

type CostSum struct {
	*cardIntProperty
}

func NewCostSum(i int) *CostSum {
	return &CostSum{&cardIntProperty{&cardProperty{}, i}}
}

type Health struct {
	*cardIntProperty
}

func NewHealth(i int) *Health {
	return &Health{&cardIntProperty{&cardProperty{}, i}}
}

type Speedmodifier struct {
	*cardIntProperty
}

func NewSpeedmodifier(i int) *Speedmodifier {
	return &Speedmodifier{&cardIntProperty{&cardProperty{}, i}}
}

type Tag struct {
	*cardStringProperty
}

func NewTag(s string) *Tag {
	return &Tag{&cardStringProperty{&cardProperty{}, s}}
}

type Text struct {
	*cardStringProperty
}

func NewText(s string) *Text {
	return &Text{&cardStringProperty{&cardProperty{}, s}}
}


type EntitiesOnBoard struct {
	*playerIntProperty
}

func NewEntitiesOnBoard(i int) *EntitiesOnBoard {
	return &EntitiesOnBoard{&playerIntProperty{&playerProperty{}, i}}
}

type DeckSize struct {
	*playerIntProperty
}

func NewDeckSize(i int) *DeckSize {
	return &DeckSize{&playerIntProperty{&playerProperty{}, i}}
}

type DustpileSize struct {
	*playerIntProperty
}

func NewDustpileSize(i int) *DustpileSize {
	return &DustpileSize{&playerIntProperty{&playerProperty{}, i}}
}

type FieldsOnBoard struct {
	*playerIntProperty
}

func NewFieldsOnBoard(i int) *FieldsOnBoard {
	return &FieldsOnBoard{&playerIntProperty{&playerProperty{}, i}}
}

type HandSize struct {
	*playerIntProperty
}

func NewHandSize(i int) *HandSize {
	return &HandSize{&playerIntProperty{&playerProperty{}, i}}
}
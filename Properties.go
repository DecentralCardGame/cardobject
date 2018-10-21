package cardobject

type Attack struct {
	*cardIntProperty
}

func NewAttack(i int) *Attack{
	return &Attack{&cardIntProperty{&cardProperty{}, i}}
}

type HandSize struct {
	*playerIntProperty
}

func NewHandsize(i int) *HandSize {
	return &HandSize{&playerIntProperty{&playerProperty{}, i}}
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
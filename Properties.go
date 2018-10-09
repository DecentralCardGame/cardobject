package cardobject

type Damage struct {
	*cardIntProperty
}

func NewDamage(i int) *Damage{
	return &Damage{&cardIntProperty{&cardProperty{}, i}}
}

type HandSize struct {
	*playerIntProperty
}

func NewHandsize(i int) *HandSize {
	return &HandSize{&playerIntProperty{&playerProperty{}, i}}
}

type Life struct {
	*cardIntProperty
}

func NewLife(i int) *Life {
	return &Life{&cardIntProperty{&cardProperty{}, i}}
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
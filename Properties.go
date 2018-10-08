package cardobject

type Damage struct{
	*cardIntProperty
}

func NewDamage(i int) *Damage{
	return &Damage{&cardIntProperty{&cardProperty{}, i}}
}

type Life struct{
	*cardIntProperty
}

func NewLife(i int) *Life{
	return &Life{&cardIntProperty{&cardProperty{}, i}}
}

type Speedmodifier struct{
	*cardIntProperty
}

type Tag struct{
	*cardStringProperty
}

type Text struct{
	*cardStringProperty
}

func NewText(s string) *Text{
	return &Text{&cardStringProperty{&cardProperty{}, s}}
}
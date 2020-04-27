package cardobject

type intCondition struct {
	IntProperty   string
	IntComparator string
	IntValue      int
}

type stringCondition struct {
	StringProperty   string
	StringComparator string
	StringValue      string
}

type tagCondition struct {
	TagComparator string
	TagValue      string
}

type conditionAttributes struct {
	IntCondition    *intCondition    `json:",omitempty"`
	StringCondition *stringCondition `json:",omitempty"`
	TagCondition    *tagCondition    `json:",omitempty"`
}

type actionCondition struct {
	conditionAttributes
}

type entityCondition struct {
	conditionAttributes
}

type placeCondition struct {
	conditionAttributes
}

type playerCondition struct {
	conditionAttributes
}

type cardCondition struct {
	ActionCondition *actionCondition `json:",omitempty"`
	EntityCondition *entityCondition `json:",omitempty"`
	PlaceCondition  *placeCondition  `json:",omitempty"`
}

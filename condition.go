package cardobject

type intCondition struct {
	IntProperty string
	IntComparator string
	IntValue int
}

type stringCondition struct {
	StringProperty string
	StringComparator string
	StringValue string
}

type conditionAttributes struct {
	IntCondition *intCondition `json:",omitempty"`
	StringCondition *stringCondition `json:",omitempty"`
}

type actionCondition struct {
	conditionAttributes
}

type entityCondition struct {
	conditionAttributes
}

type fieldCondition struct {
	conditionAttributes
}

type playerCondition struct {
	conditionAttributes
}

type cardCondition struct {
	ActionCondition *actionCondition `json:",omitempty"`
	EntityCondition *entityCondition `json:",omitempty"`
	FieldCondition *fieldCondition `json:",omitempty"`
}
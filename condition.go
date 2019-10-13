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

type conditionWrapper struct {
	IntCondition *intCondition `json:",omitempty"`
	StringCondition *stringCondition `json:",omitempty"`
}

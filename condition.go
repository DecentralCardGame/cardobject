package cardobject

type condition interface {
	GetConditionAttributes() conditionAttributes
}

type conditionAttributes struct {
	Property string
}

type intCondition struct {
	conditionAttributes
	IntValue int
	IntComparator string
}

type stringCondition struct {
	conditionAttributes
	StringValue string
	StringComparator string
}

func (ca conditionAttributes) GetConditionAttributes() conditionAttributes {
	return ca
}
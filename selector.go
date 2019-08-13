package cardobject

type selector interface {
	GetSelectorAttributes() selectorAttributes
}

type selectorAttributes struct {}

type targetingAttributes struct {
	PlayerMode string
	PlayerCondition *conditionWrapper `json:",omitempty"`
	CardMode string
	Zone string
	CardCondition *conditionWrapper `json:",omitempty"`
}

type actionSelector struct {
	selectorAttributes
	targetingAttributes
}

type entitySelector struct {
	selectorAttributes
	targetingAttributes
}

type fieldSelector struct {
	selectorAttributes
	targetingAttributes
}

type selfSelector struct {
	selectorAttributes
	Target string
}

type selectorWrapper struct {
	ActionSelector *actionSelector `json:",omitempty"`
	EntitySelector *entitySelector `json:",omitempty"`
	FieldSelector *fieldSelector `json:",omitempty"`
	SelfSelector *selfSelector `json:",omitempty"`
}

func (s selectorAttributes) GetSelectorAttributes() selectorAttributes {
	return s
}
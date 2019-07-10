package cardobject

type selector interface {
	GetSelectorAttributes() selectorAttributes
}

type selectorAttributes struct {}

type targetingAttributes struct {
	PlayerMode string
	PlayerCondition conditionWrapper
	CardMode string
	Zone string
	CardCondition conditionWrapper
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

func (s selectorAttributes) GetSelectorAttributes() selectorAttributes {
	return s
}
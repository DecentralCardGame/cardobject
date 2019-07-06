package cardobject

type selector interface {
	GetSelectorAttributes() selectorAttributes
}

type selectorAttributes struct {}

type targetingAttributes struct {
	PlayerMode string
	//PlayerCondition
	CardMode string
	Zone string
}

type actionSelector struct {
	selectorAttributes
	targetingAttributes
	//actionCondition
}

type entitySelector struct {
	selectorAttributes
	targetingAttributes
	//actionCondition
}

type fieldSelector struct {
	selectorAttributes
	targetingAttributes
	//actionCondition
}

type selfSelector struct {
	selectorAttributes
	Target string
}

func (s selectorAttributes) GetSelectorAttributes() selectorAttributes {
	return s
}
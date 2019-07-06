package cardobject

type selector interface {
	GetSelectorAttributes() selectorAttributes
}

type selectorAttributes struct {
	PlayerMode string
	//PlayerCondition
	CardMode string
	Zone string
}

type actionSelector struct {
	*selectorAttributes
	//actionCondition
}

type entitySelector struct {
	*selectorAttributes
	//actionCondition
}

type fieldSelector struct {
	*selectorAttributes
	//actionCondition
}

func (s selectorAttributes) GetSelectorAttributes() selectorAttributes {
	return s
}
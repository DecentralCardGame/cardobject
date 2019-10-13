package cardobject

type selectorAttributes struct {
	PlayerMode string
	PlayerCondition *conditionWrapper `json:",omitempty"`
	CardMode string
	CardCondition *conditionWrapper `json:",omitempty"`
	Zone string
}

type actionSelector struct {
	selectorAttributes
}

type entitySelector struct {
	selectorAttributes
}

type fieldSelector struct {
	selectorAttributes
}

type selfSelector struct {
	Target string
}
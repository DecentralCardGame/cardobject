package cardobject

type selectorAttributes struct {
	PlayerMode string
	PlayerCondition *condition `json:",omitempty"`
	CardMode string
	CardCondition *condition `json:",omitempty"`
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
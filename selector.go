package cardobject

type selectorAttributes struct {
	PlayerMode string
	PlayerCondition *playerCondition `json:",omitempty"`
	Zone string
	CardMode string
}

type actionSelector struct {
	selectorAttributes
	CardCondition *actionCondition `json:",omitempty"`
}

type entitySelector struct {
	selectorAttributes
	CardCondition *entityCondition `json:",omitempty"`
}

type locationSelector struct {
	selectorAttributes
	CardCondition *locationCondition `json:",omitempty"`
}

type selfSelector struct {
	Target string
}

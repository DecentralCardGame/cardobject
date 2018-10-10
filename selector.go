package cardobject

type Selector interface {
	IsSelector() bool
}

type basicSelector struct {
	playerMode SelectorMode
	playerCondition PlayerCondition
	cardMode SelectorMode
	cardCondition CardCondition
	location CardLocation
}

type selectorModes int

const (
	ALL selectorModes = iota
	TARGET
	OPPONENTCHOICE
)

type SelectorMode interface {
	SelectorModes() selectorModes
}

func (s selectorModes) SelectorModes() selectorModes {
	return s
}

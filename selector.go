package cardobject

type Selector struct {
	playerMode SelectorMode
	zone Zone
	playerCondition Condition
	cardMode SelectorMode
	cardCondition Condition
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

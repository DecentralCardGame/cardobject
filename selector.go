package cardobject

type Selector struct {
	mode SelectorMode
}

type PlayerSelector struct {
	*Selector
	//property + wert
}

type CardSelector struct {
	*Selector
	//property + wert
}

type selectorModes int

const (
	ALL selectorModes = iota
	RANDOM
	TARGET
	OPPONENTCHOICE
)

type SelectorMode interface {
	SelectorModes() selectorModes
}

func (s selectorModes) SelectorModes() selectorModes {
	return s
}
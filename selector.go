package cardobject

type Selector interface {
	IsSelector() bool
}

type SafeSelector interface {
	IsSafe() bool
}

func NewBasicSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, l CardLocation) Selector {
	return &basicSelector{pm, pc, cm, cc, l}
}

func NewSafeSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, dl DynamicCardLocation) SafeSelector{
	bs := NewBasicSelector(pm, pc, cm, cc, dl)
	return bs.(SafeSelector)
}


type basicSelector struct {
	playerMode SelectorMode
	playerCondition PlayerCondition
	cardMode SelectorMode
	cardCondition CardCondition
	location CardLocation
}

type safeSelector struct {
	*basicSelector
}

type simpleSelectorID int 

const (
	THIS simpleSelectorID = iota
)

func (bs *basicSelector) IsSelector() bool {
	return true
}

func (ss *simpleSelectorID) IsSelector() bool {
	return true
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

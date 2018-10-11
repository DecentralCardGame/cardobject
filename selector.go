package cardobject

type Selector interface {
	IsSelector() bool
}

type SafeSelector interface {
	IsSafe() bool
}

func NewBasicSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, z Zone) Selector {
	return &basicSelector{pm, pc, cm, cc, z}
}

func NewSafeSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, dz DynamicZone) SafeSelector{
	bs := NewBasicSelector(pm, pc, cm, cc, dz)
	return bs.(SafeSelector)
}


type basicSelector struct {
	playerMode SelectorMode
	playerCondition PlayerCondition
	cardMode SelectorMode
	cardCondition CardCondition
	zone Zone
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

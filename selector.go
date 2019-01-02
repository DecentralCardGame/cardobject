package cardobject

type CardSelector interface {
	IsCardSelector() bool
}

func NewBasicSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, z Zone) *CardSelectorCond {
	return &CardSelectorCond{pm, pc, cm, cc, z}
}

func NewSafeSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, dz DynamicZone) *CardSelectorSafeCond {
	return &CardSelectorSafeCond{&CardSelectorCond{pm, pc, cm, cc, dz}, dz}
}


type CardSelectorCond struct {
	playerMode SelectorMode
	playerCondition PlayerCondition
	cardMode SelectorMode
	cardCondition CardCondition
	zone Zone
}

type CardSelectorSafeCond struct {
	*CardSelectorCond
	safeZone DynamicZone
}

type simpleSelectorID int 

const (
	THIS simpleSelectorID = iota
)


func (bs *CardSelectorCond) IsCardSelector() bool {
	return true
}

func (ss *simpleSelectorID) IsCardSelector() bool {
	return true
}

func (bs *CardSelectorCond) PlayerSelectorMode() SelectorMode {
	return bs.playerMode
}

func (bs *CardSelectorCond) PlayerCondition() PlayerCondition {
	return bs.playerCondition
}

func (bs *CardSelectorCond) CardSelectorMode() SelectorMode {
	return bs.cardMode
}

func (bs *CardSelectorCond) CardCondition() CardCondition {
	return bs.cardCondition
}

func (bs *CardSelectorCond) Zone() Zone {
	return bs.zone
}

func (ss *CardSelectorSafeCond) DynamicZone() DynamicZone {
	return ss.safeZone
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

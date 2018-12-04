package cardobject

type CardSelector interface {
	IsCardSelector() bool
}

type CardSelectorCond interface {
	CardSelector
	PlayerSelectorMode() SelectorMode
	PlayerCondition() PlayerCondition
	CardSelectorMode() SelectorMode
	CardCondition() CardCondition
	Zone() Zone
}

type CardSelectorSafeCond interface {
	CardSelector
	DynamicZone() DynamicZone
}

func NewBasicSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, z Zone) CardSelectorCond {
	return &basicSelector{pm, pc, cm, cc, z}
}

func NewSafeSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, dz DynamicZone) CardSelectorSafeCond {
	return &safeSelector{&basicSelector{pm, pc, cm, cc, dz}, dz}
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
	safeZone DynamicZone
}

type simpleSelectorID int 

const (
	THIS simpleSelectorID = iota
)


func (bs *basicSelector) IsCardSelector() bool {
	return true
}

func (ss *simpleSelectorID) IsCardSelector() bool {
	return true
}

func (bs *basicSelector) PlayerSelectorMode() SelectorMode {
	return bs.playerMode
}

func (bs *basicSelector) PlayerCondition() PlayerCondition {
	return bs.playerCondition
}

func (bs *basicSelector) CardSelectorMode() SelectorMode {
	return bs.cardMode
}

func (bs *basicSelector) CardCondition() CardCondition {
	return bs.cardCondition
}

func (bs *basicSelector) Zone() Zone {
	return bs.zone
}

func (ss *safeSelector) DynamicZone() DynamicZone {
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

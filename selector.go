package cardobject

type CardSelector interface {
	IsCardSelector() bool
}

type CardSelectorCond interface {
	CardSelector
	GetPlayerSelectorMode() SelectorMode
	GetPlayerCondition() PlayerCondition
	GetCardSelectorMode() SelectorMode
	GetCardCondition() CardCondition
	GetZone() Zone
}

type CardSelectorSafeCond interface {
	CardSelector
	GetDynamicZone() DynamicZone
}

func NewBasicSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, z Zone) CardSelectorCond {
	return &basicSelector{pm, pc, cm, cc, z}
}

func NewSafeSelector(pm SelectorMode, pc PlayerCondition, cm SelectorMode, cc CardCondition, dz DynamicZone) CardSelectorSafeCond {
	return &safeSelector{&basicSelector{pm, pc, cm, cc, dz}, dz}
}


type basicSelector struct {
	PlayerMode SelectorMode
	PlayerCondition PlayerCondition
	CardMode SelectorMode
	CardCondition CardCondition
	Zone Zone
}

type safeSelector struct {
	*basicSelector
	SafeZone DynamicZone
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

func (bs *basicSelector) GetPlayerSelectorMode() SelectorMode {
	return bs.PlayerMode
}

func (bs *basicSelector) GetPlayerCondition() PlayerCondition {
	return bs.PlayerCondition
}

func (bs *basicSelector) GetCardSelectorMode() SelectorMode {
	return bs.CardMode
}

func (bs *basicSelector) GetCardCondition() CardCondition {
	return bs.CardCondition
}

func (bs *basicSelector) GetZone() Zone {
	return bs.Zone
}

func (ss *safeSelector) GetDynamicZone() DynamicZone {
	return ss.SafeZone
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

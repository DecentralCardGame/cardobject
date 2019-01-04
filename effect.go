package cardobject

type Effect interface {
	GetZoneChanges() []ZoneChange
	GetManipulations() []Manipulation
	GetProduction() []Ressource
}

type Manipulation interface {
	GetCardSelector() CardSelector
	GetManipulation() Inserter
	GetTargetPropertyId() PropertyId
}

type IntManipulation interface {
	Manipulation
	GetIntManipulation() IntInserter
	GetTargetIntPropertyId() IntPropertyId
}

type StringManipulation interface {
	Manipulation
	GetStringManipulation() StringInserter
	GetTargetStringPropertyId() StringPropertyId
}

func NewEffect(zc []ZoneChange, m []Manipulation, p []Ressource) Effect {
	return &effect{zc, m, p}	
}

func NewIntManipulation(cs CardSelector, v IntInserter, p IntPropertyId) IntManipulation {
	return &intManipulation{&manipulation{cs}, v, p}
}

func NewStringManipulation(cs CardSelector, v StringInserter, p StringPropertyId) StringManipulation {
	return &stringManipulation{&manipulation{cs}, v, p}
}

type effect struct {
	ZoneChanges []ZoneChange
	Manipulations []Manipulation
	Production []Ressource
}

type manipulation struct {
	selector CardSelector
}

type intManipulation struct {
	*manipulation
	val IntInserter
	prop IntPropertyId
}

type stringManipulation struct {
	*manipulation
	val StringInserter
	prop StringPropertyId
}

func (e *effect) GetZoneChanges() []ZoneChange {
	return e.ZoneChanges
}

func (e *effect) GetManipulations() []Manipulation {
	return e.Manipulations
}

func (e *effect) GetProduction() []Ressource {
	return e.Production
}

func (m *manipulation) GetCardSelector() CardSelector {
	return m.selector
}

func (im *intManipulation) GetManipulation() Inserter {
	return im.GetIntManipulation()
}

func (sm *stringManipulation) GetManipulation() Inserter {
	return sm.GetStringManipulation()
}

func (im *intManipulation) GetTargetPropertyId() PropertyId {
	return im.GetTargetIntPropertyId()
}

func (sm *stringManipulation) GetTargetPropertyId() PropertyId {
	return sm.GetTargetStringPropertyId()
}

func (im *intManipulation) GetIntManipulation() IntInserter {
	return im.val
}

func (sm *stringManipulation) GetStringManipulation() StringInserter {
	return sm.val
}

func (im *intManipulation) GetTargetIntPropertyId() IntPropertyId {
	return im.prop
}

func (sm *stringManipulation) GetTargetStringPropertyId() StringPropertyId {
	return sm.prop
}


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
	GetArithOperator() ArithOperator
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

func NewIntManipulation(cs CardSelector, v IntInserter, a ArithOperator, p IntPropertyId) IntManipulation {
	return &intManipulation{&manipulation{cs}, v, a, p}
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
	Selector CardSelector
}

type intManipulation struct {
	*manipulation
	Val IntInserter
	Operator ArithOperator
	Prop IntPropertyId
}

type stringManipulation struct {
	*manipulation
	Val StringInserter
	Prop StringPropertyId
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
	return m.Selector
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
	return im.Val
}

func (sm *stringManipulation) GetStringManipulation() StringInserter {
	return sm.Val
}

func (im *intManipulation) GetTargetIntPropertyId() IntPropertyId {
	return im.Prop
}

func (sm *stringManipulation) GetTargetStringPropertyId() StringPropertyId {
	return sm.Prop
}

func (im *intManipulation) GetArithOperator() ArithOperator {
	return im.Operator
}


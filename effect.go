package cardobject

type Effect struct {
	zoneChanges []ZoneChange
	manipulations []Manipulations
	production []Ressource
}

type Manipulations struct {
	selector Selector
	value int
	property string
}
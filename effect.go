package cardobject

type Effect struct {
	zoneChanges []ZoneChange
	manipulations []Manipulations
}

type Manipulations struct {
	selector Selector
	value int
	property string
}
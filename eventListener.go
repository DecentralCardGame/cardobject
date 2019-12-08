package cardobject

type productionEventListener struct {
	RessourceType string
}

type timeEventListener struct {
	TimeEvent string
}

type manipulationEventListener struct {
	IntManipulationEventListener *manipulationEventBasics `json:",omitempty"`
	StringManipulationEventListener *manipulationEventBasics `json:",omitempty"`
}

type manipulationEventBasics struct {
	Property string
	ChangeMode string
	CardCondition *cardCondition `json:",omitempty"`
}

type zoneChangeEventListener struct {
	Source string
	Destination string
	CardCondition *cardCondition `json:",omitempty"`
}

type eventListener struct {
	ProductionEventListener *productionEventListener `json:",omitempty"`
	TimeEventListener *timeEventListener `json:",omitempty"`
	ManipulationEventListener *manipulationEventListener `json:",omitempty"`
	ZoneChangeEventListener *zoneChangeEventListener `json:",omitempty"`
}
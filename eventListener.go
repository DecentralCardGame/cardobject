package cardobject

type timeEventListener struct {
	TimeEvent string
}

type manipulationEventListener struct {
	Property string
}

type zoneChangeEventListener struct {
	Source string
	Destination string
}

type eventListener struct {
	TimeEventListener *timeEventListener `json:",omitempty"`
	ManipulationEventListener *manipulationEventListener `json:",omitempty"`
	ZoneChangeEventListener *zoneChangeEventListener `json:",omitempty"`
}
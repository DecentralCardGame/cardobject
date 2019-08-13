package cardobject

type eventListener interface {
	GetEventListenerAttributes() eventListenerAttributes
}

type eventListenerAttributes struct {}

type timeEventListener struct {
	eventListenerAttributes
	TimeEvent string
}

type manipulationEventListener struct {
	eventListenerAttributes
	Property string
}

type zoneChangeEventListener struct {
	eventListenerAttributes
	Source string
	Destination string
}

type eventListenerWrapper struct {
	TimeEventListener *timeEventListener `json:",omitempty"`
	ManipulationEventListener *manipulationEventListener `json:",omitempty"`
	ZoneChangeEventListener *zoneChangeEventListener `json:",omitempty"`
}

func (ev eventListenerAttributes) GetEventListenerAttributes() eventListenerAttributes {
	return ev
}
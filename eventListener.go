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

func (ev eventListenerAttributes) GetEventListenerAttributes() eventListenerAttributes {
	return ev
}
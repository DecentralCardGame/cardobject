package cardobject

type EventListener interface {
	IsEventListener() bool
}

type ReturningEventListener interface {
	EventListener
	ReturnsCardPointerArray() bool
}

type TimeEventListener interface {
	EventListener
	TimeEvent() TimeEvent
}

type ManipulationEventListener interface {
	ReturningEventListener
	PropertyId() CardPropertyId
}

type ZoneChangeEvenListener interface {
	ReturningEventListener
	Source() DynamicZone
	Destination() Zone
}

func NewTimeEventListener(te TimeEvent) TimeEventListener {
	return &timeEventListener{&eventListener{}, te}
}

func NewManipulationEventListener(pi CardPropertyId) ManipulationEventListener {
	return &manipulationEventListener{&eventListener{}, pi}
}

func NewZoneChangeEvenListener(d DynamicZone, s Zone) ZoneChangeEvenListener {
	return &zoneChangeEventListener{&eventListener{}, d, s}
}


type eventListener struct {}

type timeEventListener struct {
	*eventListener
	event TimeEvent
}

type manipulationEventListener struct {
	*eventListener
	propertyId CardPropertyId
}

type zoneChangeEventListener struct {
	*eventListener
	source DynamicZone
	destination Zone
}


func (el *eventListener) IsEventListener() bool {
	return true
}

func (mel *manipulationEventListener) ReturnsCardPointerArray() bool {
	return true
}

func (zel *zoneChangeEventListener) ReturnsCardPointerArray() bool {
	return true
}

func (tel *timeEventListener) TimeEvent() TimeEvent {
	return tel.event
}

func (mel *manipulationEventListener) PropertyId() CardPropertyId {
	return mel.propertyId
}

func (zel *zoneChangeEventListener) Source() DynamicZone {
	return zel.source
}

func (zel *zoneChangeEventListener) Destination() Zone {
	return zel.destination
}
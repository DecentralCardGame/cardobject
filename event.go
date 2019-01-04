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
	GetTimeEvent() TimeEvent
}

type ManipulationEventListener interface {
	ReturningEventListener
	GetPropertyId() CardPropertyId
}

type ZoneChangeEventListener interface {
	ReturningEventListener
	GetSource() DynamicZone
	GetDestination() Zone
}

func NewTimeEventListener(te TimeEvent) TimeEventListener {
	return &timeEventListener{&eventListener{}, te}
}

func NewManipulationEventListener(pi CardPropertyId) ManipulationEventListener {
	return &manipulationEventListener{&eventListener{}, pi}
}

func NewZoneChangeEventListener(d DynamicZone, s Zone) ZoneChangeEventListener {
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

func (tel *timeEventListener) GetTimeEvent() TimeEvent {
	return tel.event
}

func (mel *manipulationEventListener) GetPropertyId() CardPropertyId {
	return mel.propertyId
}

func (zel *zoneChangeEventListener) GetSource() DynamicZone {
	return zel.source
}

func (zel *zoneChangeEventListener) GetDestination() Zone {
	return zel.destination
}
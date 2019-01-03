package cardobject

type EventListener interface {
	IsEventListener() bool
}

type ReturningEventListener interface {
	EventListener
	ReturnsCardPointerArray() bool
}

func NewTimeEventListener(te TimeEvent) *TimeEventListener {
	return &TimeEventListener{&eventListener{}, te}
}

func NewManipulationEventListener(pi CardPropertyId) *ManipulationEventListener {
	return &ManipulationEventListener{&eventListener{}, pi}
}

func NewZoneChangeEvenListener(d DynamicZone, s Zone) *ZoneChangeEventListener {
	return &ZoneChangeEventListener{&eventListener{}, d, s}
}


type eventListener struct {}

type TimeEventListener struct {
	*eventListener
	event TimeEvent
}

type ManipulationEventListener struct {
	*eventListener
	propertyId CardPropertyId
}

type ZoneChangeEventListener struct {
	*eventListener
	source DynamicZone
	destination Zone
}


func (el *eventListener) IsEventListener() bool {
	return true
}

func (mel *ManipulationEventListener) ReturnsCardPointerArray() bool {
	return true
}

func (zel *ZoneChangeEventListener) ReturnsCardPointerArray() bool {
	return true
}

func (tel *TimeEventListener) TimeEvent() TimeEvent {
	return tel.event
}

func (mel *ManipulationEventListener) PropertyId() CardPropertyId {
	return mel.propertyId
}

func (zel *ZoneChangeEventListener) Source() DynamicZone {
	return zel.source
}

func (zel *ZoneChangeEventListener) Destination() Zone {
	return zel.destination
}
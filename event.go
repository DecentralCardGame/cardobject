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
	Event TimeEvent
}

type manipulationEventListener struct {
	*eventListener
	PropertyId CardPropertyId
}

type zoneChangeEventListener struct {
	*eventListener
	Source DynamicZone
	Destination Zone
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
	return tel.Event
}

func (mel *manipulationEventListener) GetPropertyId() CardPropertyId {
	return mel.PropertyId
}

func (zel *zoneChangeEventListener) GetSource() DynamicZone {
	return zel.Source
}

func (zel *zoneChangeEventListener) GetDestination() Zone {
	return zel.Destination
}
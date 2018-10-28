package cardobject

type EventListener interface {
	IsEventListener() bool
}

type ReturningEventListener interface {
	EventListener
	ReturnsCardPointer() bool
}

type TimeEventListener interface {
	EventListener
}

type ManipulationEventListener interface {
	ReturningEventListener
}

type ZoneChangeEvenListener interface {
	ReturningEventListener
}

func NewTimeEventListener(s bool, e bool) TimeEventListener {
	return &timeEventListener{&eventListener{}, s, e}
}

func NewManipulationEventListener(p CardProperty) ManipulationEventListener {
	return &manipulationEventListener{&eventListener{}, p}
}

func NewZoneChangeEvenListener(s Zone, d DynamicZone) ZoneChangeEvenListener {
	return &zoneChangeEventListener{&eventListener{}, s, d}
}


type eventListener struct {}

type timeEventListener struct {
	*eventListener
	startOfTick bool
	endOfTick bool
}

type manipulationEventListener struct {
	*eventListener
	property CardProperty
}

type zoneChangeEventListener struct {
	*eventListener
	source Zone
	destination DynamicZone
}


func (el *eventListener) IsEventListener() bool {
	return true
}

func (mel *manipulationEventListener) ReturnsCardPointer() bool {
	return true
}

func (zel *zoneChangeEventListener) ReturnsCardPointer() bool {
	return true
}


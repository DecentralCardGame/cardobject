package cardobject

type EventListener interface {
	CheckEvent(Event) bool
}

type eventListener struct {

}

type ReturningEventListener interface {
	EventListener
	GetCard(Event) *Card
}

type returningEventListener struct {
	*eventListener
}

func (el *eventListener) CheckEvent(Event) bool {
	return false
}

func (el *eventListener) GetCard(Event) *Card {
	return &Card{}
}

type ZoneChangeListener struct {
	*returningEventListener
	Origin Zone
	Destination Zone
	Cardtype Card
}

type StateChangeListener struct {
	*returningEventListener
	State CardProperty
	Cardtype Card
}

type TikChangeListener struct {
	*eventListener
}


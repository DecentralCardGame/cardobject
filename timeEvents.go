package cardobject

const (
	TICKSTART timeEvent = iota
	TICKEND
)

type TimeEvent interface {
	TimeEvent() timeEvent
}


type timeEvent int

func (te timeEvent) TimeEvent() timeEvent {
	return te
}

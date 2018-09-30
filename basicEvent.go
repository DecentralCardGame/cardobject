package cardobject

type events int

const (
	TODO events = iota
)

type Event interface {
	Events() events
}

func (r events) Events() events {
	return r
}

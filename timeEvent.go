package cardobject

import "github.com/DecentralCardGame/jsonschema"

var timeEvents []string = []string{"TICKSTART", "COMBAT"}

type timeEvent struct{ *jsonschema.BasicEnum }

func (t timeEvent) GetEnumValues() []string {
	return timeEvents
}

package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var timeEvents []string = []string{"TICKSTART", "COMBAT"}

type TimeEvent jsonschema.BasicEnum

func (t TimeEvent) Validate() error {
	return t.ValidateEnum()
}

func (t TimeEvent) ValidateEnum() error {
	values := t.GetEnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("")
}

func (t TimeEvent) GetEnumValues() []string {
	return timeEvents
}

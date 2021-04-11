package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

var timeEvents []string = []string{"TICKSTART", "COMBAT"}

type TimeEvent jsonschema.BasicEnum

func (t TimeEvent) Validate() error {
	return t.ValidateEnum()
}

func (t TimeEvent) ValidateEnum() error {
	values := t.EnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("TimeEvent must be one of: " + strings.Join(timeEvents, ","))
}

func (t TimeEvent) EnumValues() []string {
	return timeEvents
}

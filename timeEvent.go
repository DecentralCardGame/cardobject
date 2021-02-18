package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var timeEvents []string = []string{"TICKSTART", "COMBAT"}

type timeEvent jsonschema.BasicEnum

func (t timeEvent) Validate() error {
	return t.ValidateEnum()
}

func (t timeEvent) ValidateEnum() error {
	values := t.GetEnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("")
}

func (t timeEvent) GetEnumValues() []string {
	return timeEvents
}

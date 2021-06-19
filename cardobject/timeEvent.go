package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//Combat TimeEvent
const Combat = "COMBAT"

//Tickstart TimeEvent
const Tickstart = "TICKSTART"

var timeEvents []string = []string{Combat, Tickstart}

type TimeEvent jsonschema.BasicEnum

func (t TimeEvent) Validate(r jsonschema.RootElement) error {
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

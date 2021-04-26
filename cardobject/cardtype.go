package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

var possibleCardTypes []string = []string{
	"ACTION",
	"ENTITY",
	"PLACE",
	"HQ"}

type CardType jsonschema.BasicEnum

func (c CardType) Validate() error {
	return c.ValidateEnum()
}

func (c CardType) ValidateEnum() error {
	values := c.EnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("Tag must be one of: " + strings.Join(c.EnumValues(), ","))
}

func (c CardType) EnumValues() []string {
	return possibleCardTypes
}

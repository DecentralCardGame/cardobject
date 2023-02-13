package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

const ACTIONTYPE string = "ACTION"
const ENTITYTYPE string = "ENTITY"
const PLACETYPE string = "PLACE"
const HEADQUARTERTYPE string = "HQ"

var possibleCardTypes []string = []string{
	ACTIONTYPE,
	ENTITYTYPE,
	PLACETYPE,
	HEADQUARTERTYPE}

type CardType jsonschema.BasicEnum

func (c CardType) ValidateType(r jsonschema.RootElement) error {
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

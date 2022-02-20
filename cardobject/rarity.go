package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

var possibleRarities []string = []string{
	"COMMON",
	"UNCOMMON",
	"RARE"}

type Rarity jsonschema.BasicEnum

func (ra Rarity) ValidateType(r jsonschema.RootElement) error {
	values := ra.EnumValues()
	for _, v := range values {
		if v == string(ra) {
			return nil
		}
	}
	return errors.New("Rarity must be one of: " + strings.Join(ra.EnumValues(), ","))
}

func (r Rarity) EnumValues() []string {
	return possibleRarities
}

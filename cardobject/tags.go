package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

var possibleTags []string = []string{
	"ANIMAL",
	"BOT",
	"DWARF",
	"ENGINEER",
	"EQUIPMENT",
	"FARM",
	"FIRE",
	"HUMAN",
	"KNIGHT",
	"MAGIC",
	"MILITANT",
	"PRIMITIVE",
	"RANGE",
	"SPIRITUAL",
	"TACTIC",
	"TECHNOCRAT"}

type Tags []Tag

func (t Tags) Validate(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range t {
		err := v.Validate(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (t Tags) MinMaxItems() (int, int) {
	return 1, 3
}

func (t Tags) ItemName() string {
	return jsonschema.ItemNameFromArray(t)
}

type Tag jsonschema.BasicEnum

func (t Tag) Validate(r jsonschema.RootElement) error {
	values := t.EnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("Tag must be one of: " + strings.Join(t.EnumValues(), ","))
}

func (t Tag) EnumValues() []string {
	return possibleTags
}

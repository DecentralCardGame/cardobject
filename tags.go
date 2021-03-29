package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/jsonschema"
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

func (t Tags) Validate() error {
	return t.ValidateArray()
}

func (t Tags) ValidateArray() error {
	errorRange := []error{}
	for _, v := range t {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (t Tags) GetMinMaxItems() (int, int) {
	return 1, 3
}

func (t Tags) GetItemName() string {
	return jsonschema.GetItemNameFromArray(t)
}

type Tag jsonschema.BasicEnum

func (t Tag) Validate() error {
	return t.ValidateEnum()
}

func (t Tag) ValidateEnum() error {
	values := t.GetEnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("Tag must be one of: " + strings.Join(possibleTags, ","))
}

func (t Tag) GetEnumValues() []string {
	return possibleTags
}

package cardobject

import "github.com/DecentralCardGame/jsonschema"

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

type tags []tag

func (t tags) Validate() error {
	return t.ValidateArray()
}

func (t tags) ValidateArray() error {
	return nil
}

func (t tags) GetMinMaxItems() (int, int) {
	return 1, 3
}

type tag struct{ *jsonschema.BasicEnum }

func (t tag) GetEnumValues() []string {
	return possibleTags
}

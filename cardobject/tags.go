package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

var possibleTags []string = []string{
	"AIR",
	"ALCHEMIST",
	"ANIMAL",
	"ARTIFACT",
	"ASSASINE",
	"BEAST",
	"BELIEVER",
	"BOT",
	"BUILDING",
	"COLOSSUS",
	"DESASTER",
	"DOMESTIC",
	"DRAGON",
	"EARTH",
	"EPIC",
	"EVENT",
	"FACTORY",
	"FIRE",
	"HUMAN",
	"KNOWLEDGE",
	"LANDSCAPE",
	"PLANT",
	"REINFORCEMENT",
	"WIZARD",
	"SHRINE",
	"SPIRIT",
	"TACTIC",
	"TECHNOCRAT",
	"UNDEAD",
	"VEHICLE",
	"WARRIOR",
	"WATER",
	"WEAPON",
	"WORKER"}

type Tags []Tag

func (t Tags) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range t {
		err := v.ValidateType(r)
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

func (t Tag) ValidateType(r jsonschema.RootElement) error {
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

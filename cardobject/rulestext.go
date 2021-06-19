package cardobject

import (
	"errors"
	"strconv"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

const maxRulesTextLength int = 1000
const minRulesTextLength int = 0

type RulesTexts []RulesText

func (rt RulesTexts) Validate(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range rt {
		err := v.Validate(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (r RulesTexts) MinMaxItems() (int, int) {
	return 0, 3
}

func (r RulesTexts) ItemName() string {
	return jsonschema.ItemNameFromArray(r)
}

type RulesText jsonschema.BasicString

func (rt RulesText) Validate(r jsonschema.RootElement) error {
	minLength, maxLength := rt.MinMaxLength()
	length := len(string(rt))
	if length < minLength || length > maxLength {
		return errors.New("RulesText must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (r RulesText) MinMaxLength() (int, int) {
	return minRulesTextLength, maxRulesTextLength
}

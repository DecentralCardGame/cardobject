package cardobject

import (
	"errors"
	"strconv"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

const maxRulesTextLength int = 1000
const minRulesTextLength int = 0

type RulesTexts []RulesText

func (r RulesTexts) Validate() error {
	return r.ValidateArray()
}

func (r RulesTexts) ValidateArray() error {
	errorRange := []error{}
	for _, v := range r {
		err := v.Validate()
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

func (r RulesText) Validate() error {
	return r.ValidateString()
}

func (r RulesText) ValidateString() error {
	minLength, maxLength := r.MinMaxLength()
	length := len(string(r))
	if length < minLength || length > maxLength {
		return errors.New("RulesText must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (r RulesText) MinMaxLength() (int, int) {
	return minRulesTextLength, maxRulesTextLength
}

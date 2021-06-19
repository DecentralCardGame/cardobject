package cardobject

import (
	"errors"
	"strconv"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

const maxKeywordLength int = 10000
const minKeywordLength int = 0

//Anthem Keyword
const Anthem = "ANTHEM"

//Arm Keyword
const Arm = "ARM"

//Arrival Keyword
const Arrival = "ARRIVAL"

//Battlecry Keyword
const Battlecry = "BATTLECRY"

//Harm Keyword
const Harm = "HARM"

//Kill Keyword
const Kill = "KILL"

//OnConstruction Keyword
const OnConstruction = "ONCONSTRUCTION"

//OnDeath Keyword
const OnDeath = "ONDEATH"

//OnSpawn Keyword
const OnSpawn = "ONSPAWN"

//Pay Keyword
const Pay = "PAY"

//Periodic Keyword
const Periodic = "PERIODIC"

//Produce Keyword
const Produce = "PRODUCE"

//Repair Keyword
const Repair = "REPAIR"

//Tribute Keyword
const Tribute = "TRIBUTE"

type Keywords []Keyword

func (k Keywords) Validate(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range k {
		err := v.Validate(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (k Keywords) MinMaxItems() (int, int) {
	return 1, 3
}

func (k Keywords) ItemName() string {
	return jsonschema.ItemNameFromArray(k)
}

type Keyword jsonschema.BasicString

func (k Keyword) Validate(r jsonschema.RootElement) error {
	minLength, maxLength := k.MinMaxLength()
	length := len(string(k))
	if length < minLength || length > maxLength {
		return errors.New("FlavourText must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (k Keyword) MinMaxLength() (int, int) {
	return minKeywordLength, maxKeywordLength
}

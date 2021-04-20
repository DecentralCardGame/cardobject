package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

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

var possibleKeywords []string = []string{
	Anthem, Arm, Arrival, Battlecry, Harm, Kill, OnConstruction, OnDeath, OnSpawn, Pay, Periodic, Produce, Repair, Tribute}

type Keyword jsonschema.BasicString

func (k Keyword) Validate() error {
	return k.ValidateEnum()
}

func (k Keyword) ValidateEnum() error {
	values := k.EnumValues()
	for _, v := range values {
		if v == string(k) {
			return nil
		}
	}
	return errors.New("Keyword must be one of: " + strings.Join(k.EnumValues(), ","))
}

func (k Keyword) EnumValues() []string {
	return possibleKeywords
}

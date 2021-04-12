package keywords

import (
	"reflect"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type iability interface {
	Resolve() cardobject.Ability
}

type abilities []ability

func (a abilities) Resolve() cardobject.Abilities {
	abilities := cardobject.Abilities{}
	for _, ability := range a {
		abilities = append(abilities, ability.Resolve())
	}
	return abilities
}

func (a abilities) Validate() error {
	return a.ValidateArray()
}

func (a abilities) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a abilities) MinMaxItems() (int, int) {
	return 0, 3
}

func (a abilities) ItemName() string {
	return "Ability"
}

type ability struct {
	Arrival        *arrival        `json:",omitempty"`
	Battlecry      *battlecry      `json:",omitempty"`
	OnConstruction *onConstruction `json:",omitempty"`
	OnDeath        *onDeath        `json:",omitempty"`
	OnSpawn        *onSpawn        `json:",omitempty"`
	Pay            *pay            `json:",omitempty"`
	Periodic       *periodic       `json:",omitempty"`
	Tribute        *tribute        `json:",omitempty"`
}

func (a ability) Resolve() cardobject.Ability {
	valueOfB := reflect.ValueOf(a)
	typeOfB := reflect.TypeOf(a)
	possibleImplementer := []iability{}
	for k := 0; k < valueOfB.NumField(); k++ {
		if !typeOfB.Field(k).Anonymous {
			possibleImplementer = append(possibleImplementer, valueOfB.Field(k).Interface().(iability))
		}
	}
	for _, b := range possibleImplementer {
		if !reflect.ValueOf(b).IsNil() {
			return b.Resolve()
		}
	}
	return cardobject.Ability{}
}

func (a ability) Validate() error {
	return a.ValidateInterface()
}

func (a ability) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

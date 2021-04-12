package keywords

import (
	"reflect"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ieffect interface {
	Resolve() cardobject.Effect
}

type effects []effect

func (e effects) Resolve() cardobject.Effects {
	effects := cardobject.Effects{}
	for _, effect := range e {
		effects = append(effects, effect.Resolve())
	}
	return effects
}

func (e effects) Validate() error {
	return e.ValidateArray()
}

func (e effects) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e effects) MinMaxItems() (int, int) {
	return 0, 3
}

func (e effects) ItemName() string {
	return "Effect"
}

type effect struct {
	Anthem  *anthem  `json:",omitempty"`
	Arm     *arm     `json:",omitempty"`
	Harm    *harm    `json:",omitempty"`
	Kill    *kill    `json:",omitempty"`
	Produce *produce `json:",omitempty"`
	Repair  *repair  `json:",omitempty"`
}

func (e effect) Resolve() cardobject.Effect {
	valueOfB := reflect.ValueOf(e)
	typeOfB := reflect.TypeOf(e)
	possibleImplementer := []ieffect{}
	for k := 0; k < valueOfB.NumField(); k++ {
		if !typeOfB.Field(k).Anonymous {
			possibleImplementer = append(possibleImplementer, valueOfB.Field(k).Interface().(ieffect))
		}
	}
	for _, b := range possibleImplementer {
		if !reflect.ValueOf(b).IsNil() {
			return b.Resolve()
		}
	}
	return cardobject.Effect{}
}

func (e effect) Validate() error {
	return e.ValidateInterface()
}

func (e effect) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

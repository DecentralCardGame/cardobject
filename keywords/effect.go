package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type effects []effect

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
	Anthem     *anthem     `json:",omitempty"`
	Arm        *arm        `json:",omitempty"`
	Bounce     *bounce     `json:",omitempty"`
	Burn       *burn       `json:",omitempty"`
	Discount   *discount   `json:",omitempty"`
	DrawAction *drawAction `json:",omitempty"`
	DrawEntity *drawEntity `json:",omitempty"`
	DrawPlace  *drawPlace  `json:",omitempty"`
	Grow       *grow       `json:",omitempty"`
	Harm       *harm       `json:",omitempty"`
	Heal       *heal       `json:",omitempty"`
	Kill       *kill       `json:",omitempty"`
	Insight    *insight    `json:",omitempty"`
	Mill       *mill       `json:",omitempty"`
	Produce    *produce    `json:",omitempty"`
	Reassemble *reassemble `json:",omitempty"`
	Resurrect  *resurrect  `json:",omitempty"`
	Repair     *repair     `json:",omitempty"`
	Silence    *silence    `json:",omitempty"`
	Void       *void       `json:",omitempty"`
	Withdraw   *withdraw   `json:",omitempty"`
}

func (e effect) Validate() error {
	return e.ValidateInterface()
}

func (e effect) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

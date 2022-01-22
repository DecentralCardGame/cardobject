package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type effects []effect

func (e effects) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range e {
		err := v.ValidateType(r)
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
	Ambush        *ambush        `json:",omitempty"`
	Anthem        *anthem        `json:",omitempty"`
	Arm           *arm           `json:",omitempty"`
	Armor         *armor         `json:",omitempty"`
	Bounce        *bounce        `json:",omitempty"`
	Burn          *burn          `json:",omitempty"`
	Count         *count         `json:",omitempty"`
	CountPower    *countPower    `json:",omitempty"`
	Dice          *dice          `json:",omitempty"`
	Discard       *discard       `json:",omitempty"`
	Discount      *discount      `json:",omitempty"`
	DrawAction    *drawAction    `json:",omitempty"`
	DrawEntity    *drawEntity    `json:",omitempty"`
	DrawPlace     *drawPlace     `json:",omitempty"`
	Grow          *grow          `json:",omitempty"`
	Harm          *harm          `json:",omitempty"`
	Heal          *heal          `json:",omitempty"`
	Kill          *kill          `json:",omitempty"`
	Insight       *insight       `json:",omitempty"`
	Mill          *mill          `json:",omitempty"`
	Produce       *produce       `json:",omitempty"`
	Ravage        *ravage        `json:",omitempty"`
	Reassemble    *reassemble    `json:",omitempty"`
	RecoverAction *recoverAction `json:",omitempty"`
	RecoverEntity *recoverEntity `json:",omitempty"`
	RecoverPlace  *recoverPlace  `json:",omitempty"`
	Resurrect     *resurrect     `json:",omitempty"`
	Repair        *repair        `json:",omitempty"`
	Sacrifice     *sacrifice     `json:",omitempty"`
	SelfBurn      *selfBurn      `json:",omitempty"`
	Silence       *silence       `json:",omitempty"`
	Spawn         *spawn         `json:",omitempty"`
	Strengthen    *strengthen    `json:",omitempty"`
	Tag           *tag           `json:",omitempty"`
	Train         *train         `json:",omitempty"`
	Void          *void          `json:",omitempty"`
	Withdraw      *withdraw      `json:",omitempty"`
}

func (e effect) ValidateType(r jsonschema.RootElement) error {
	implementer, err := e.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (e effect) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(e)
}

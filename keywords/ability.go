package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type abilities []ability

func (a abilities) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range a {
		err := v.ValidateType(r)
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
	Channel        *channel        `json:",omitempty"`
	DiscardPay     *discardPay     `json:",omitempty"`
	Dismantle      *dismantle      `json:",omitempty"`
	Dissolve       *dissolve       `json:",omitempty"`
	Loot           *loot           `json:",omitempty"`
	OnConstruction *onConstruction `json:",omitempty"`
	OnDeath        *onDeath        `json:",omitempty"`
	OnSpawn        *onSpawn        `json:",omitempty"`
	Pay            *pay            `json:",omitempty"`
	Periodic       *periodic       `json:",omitempty"`
	Tribute        *tribute        `json:",omitempty"`
}

func (a ability) ValidateType(r jsonschema.RootElement) error {
	implementer, err := a.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (a ability) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(a)
}

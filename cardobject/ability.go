package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type Abilities []Ability

func (a Abilities) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range a {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a Abilities) MinMaxItems() (int, int) {
	return 0, maxAbilityEffectCount
}

func (a Abilities) ItemName() string {
	return "Ability"
}

type Ability struct {
	ActivatedAbility *ActivatedAbility `json:",omitempty"`
	TriggeredAbility *TriggeredAbility `json:",omitempty"`
}

func (a Ability) ValidateType(r jsonschema.RootElement) error {
	implementer, err := a.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (a Ability) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(a)
}

type ActivatedAbility struct {
	AbilityCost *Cost
	Effects     Effects
}

func (a ActivatedAbility) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActivatedAbility) InteractionText() string {
	return "Pay §AbilityCost: §Effects \n"
}

type TriggeredAbility struct {
	Cause   *EventListener
	Cost    *Cost `json:",omitempty"`
	Effects Effects
}

func (t TriggeredAbility) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (a TriggeredAbility) InteractionText() string {
	return "§Cause , §Cost : §Effects \n"
}

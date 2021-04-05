package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type Abilities []Ability

func (a Abilities) Validate() error {
	return a.ValidateArray()
}

func (a Abilities) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
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

func (a Ability) Validate() error {
	return a.ValidateInterface()
}

func (a Ability) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type ActivatedAbility struct {
	AbilityCost *Cost
	Effects     Effects
}

func (a ActivatedAbility) Validate() error {
	return a.ValidateStruct()
}

func (a ActivatedAbility) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActivatedAbility) InteractionText() string {
	return "Pay §AbilityCost: §Effects \n"
}

type TriggeredAbility struct {
	Cause   *EventListener
	Cost    *Cost `json:",omitempty"`
	Effects Effects
}

func (t TriggeredAbility) Validate() error {
	return t.ValidateStruct()
}

func (t TriggeredAbility) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (a TriggeredAbility) InteractionText() string {
	return "§Cause , §Cost : §Effects \n"
}

package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type abilities []ability

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

func (a abilities) GetMinMaxItems() (int, int) {
	return 0, maxAbilityEffectCount
}

func (a abilities) GetItemName() string {
	return "Ability"
}

type ability struct {
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}

func (a ability) Validate() error {
	return a.ValidateInterface()
}

func (a ability) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type activatedAbility struct {
	AbilityCost *cost
	Effects     effects
}

func (a activatedAbility) Validate() error {
	return a.ValidateStruct()
}

func (a activatedAbility) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a activatedAbility) GetInteractionText() string {
	return "Pay §AbilityCost: §Effects \n"
}

type triggeredAbility struct {
	Cause   *eventListener
	Cost    *cost
	Effects effects
}

func (t triggeredAbility) Validate() error {
	return t.ValidateStruct()
}

func (t triggeredAbility) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (a triggeredAbility) GetInteractionText() string {
	return "§Cause, §Cost : §Effects \n"
}

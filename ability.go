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

type ability struct {
	*jsonschema.BasicInterface
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}

type activatedAbility struct {
	*jsonschema.BasicStruct
	AbilityCost *cost
	Effects     effects
}

func (a activatedAbility) GetInteractionText() string {
	return "Pay §AbilityCost: §Effects \n"
}

type triggeredAbility struct {
	*jsonschema.BasicStruct
	Cause   *eventListener
	Cost    *cost
	Effects effects
}

func (a triggeredAbility) GetInteractionText() string {
	return "§Cause, §Cost : §Effects \n"
}

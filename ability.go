package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type abilities []ability

func (a abilities) Validate() error {
	return a.ValidateArray()
}

func (a abilities) ValidateArray() error {
	return nil
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
	return "troll"
}

type triggeredAbility struct {
	*jsonschema.BasicStruct
	Cause   *eventListener
	Cost    *cost
	Effects effects
}

func (a triggeredAbility) GetInteractionText() string {
	return "troll"
}

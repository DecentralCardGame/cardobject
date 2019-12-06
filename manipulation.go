package cardobject

type abilityManipulationBasics struct {
	AbilityOperator string
	Ability ability
}

type effectManipulationBasics struct {
	EffectOperator string
	Effect effect
}

type intManipulationBasics struct {
	IntProperty string
	IntOperator string
	IntValue int
}

type stringManipulationBasics struct {
	StringProperty string
	StringOperator string
	StringValue string
}

type actionManipulation struct {
	ActionEffectManipulation *effectManipulationBasics `json:",omitempty"`
	ActionIntManipulation *intManipulationBasics `json:",omitempty"`
	ActionStringManipulation *stringManipulationBasics `json:",omitempty"`
}

type entityManipulation struct {
	EntityAbilityManipulation *abilityManipulationBasics `json:",omitempty"`
	EntityIntManipulation *intManipulationBasics `json:",omitempty"`
	EntityStringManipulation *stringManipulationBasics `json:",omitempty"`
}

type locationManipulation struct {
	LocationAbilityManipulation *abilityManipulationBasics `json:",omitempty"`
	LocationIntManipulation *intManipulationBasics `json:",omitempty"`
	LocationStringManipulation *stringManipulationBasics `json:",omitempty"`
}

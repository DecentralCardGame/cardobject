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

type tagManipulationBasics struct {
	TagOperator string
	TagValue string
}

type actionManipulation struct {
	ActionEffectManipulation *effectManipulationBasics `json:",omitempty"`
	ActionIntManipulation *intManipulationBasics `json:",omitempty"`
	ActionStringManipulation *stringManipulationBasics `json:",omitempty"`
	ActionTagManipulation *tagManipulationBasics `json:",omitempty"`
}

type entityManipulation struct {
	EntityAbilityManipulation *abilityManipulationBasics `json:",omitempty"`
	EntityIntManipulation *intManipulationBasics `json:",omitempty"`
	EntityStringManipulation *stringManipulationBasics `json:",omitempty"`
	EntityTagManipulation *tagManipulationBasics `json:",omitempty"`
}

type placeManipulation struct {
	PlaceAbilityManipulation *abilityManipulationBasics `json:",omitempty"`
	PlaceIntManipulation *intManipulationBasics `json:",omitempty"`
	PlaceStringManipulation *stringManipulationBasics `json:",omitempty"`
	PlaceTagManipulation *tagManipulationBasics `json:",omitempty"`
}

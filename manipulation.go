package cardobject

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
	ActionIntManipulation *intManipulationBasics `json:",omitempty"`
	ActionStringManipulation *stringManipulationBasics `json:",omitempty"`
}

type entityManipulation struct {
	EntityIntManipulation *intManipulationBasics `json:",omitempty"`
	EntityStringManipulation *stringManipulationBasics `json:",omitempty"`
}

type locationManipulation struct {
	LocationIntManipulation *intManipulationBasics `json:",omitempty"`
	LocationStringManipulation *stringManipulationBasics `json:",omitempty"`
}

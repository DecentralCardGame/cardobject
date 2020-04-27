package cardobject

type token struct {
	Name         *string `json:",omitempty"`
	Tag          *string `json:",omitempty"`
	AbilitySpeed int8
	Health       int8
	Attack       int8
}

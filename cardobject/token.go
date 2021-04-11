package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type Token struct {
	Name   CardName `json:",omitempty"`
	Attack Attack
	Health Health
	Tags   Tags `json:",omitempty"`
}

func (t Token) Validate() error {
	return t.ValidateStruct()
}

func (t Token) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t Token) InteractionText() string {
	return "§Attack / §Health token named §Name tagged §Tags"
}

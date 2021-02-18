package cardobject

import "github.com/DecentralCardGame/jsonschema"

type token struct {
	Name   cardName `json:",omitempty"`
	Attack attack
	Health health
	Tags   tags `json:",omitempty"`
}

func (t token) Validate() error {
	return t.ValidateStruct()
}

func (t token) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t token) GetInteractionText() string {
	return "§Attack / §Health token named §Name tagged §Tags"
}

package cardobject

import "github.com/DecentralCardGame/jsonschema"

type token struct {
	*jsonschema.BasicStruct
	Name   cardName `json:",omitempty"`
	Attack attack
	Health health
	Tags   tags `json:",omitempty"`
}

func (t token) GetInteractionText() string {
	return "§Attack / §Health token named §Name tagged §Tags"
}

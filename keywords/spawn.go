package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type spawn struct {
	TokenType cardobject.TokenType
	Amount    cardobject.IntValue
}

func (s spawn) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s spawn) InteractionText() string {
	return "Spawn §Amount §TokenType."
}

func (s spawn) Description() string {
	return "Spawns in a number of Tokens."
}

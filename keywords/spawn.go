package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type spawn struct {
	TokenType cardobject.TokenType
	Amount    cardobject.SimpleIntValue
}

func (s spawn) Validate() error {
	return s.ValidateStruct()
}

func (s spawn) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s spawn) InteractionText() string {
	return "Spawn §Amount §TokenType."
}

func (s spawn) Description() string {
	return "Spawns in a number of Tokens."
}

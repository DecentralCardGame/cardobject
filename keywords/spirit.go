package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type spirit struct{}

func (s spirit) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s spirit) InteractionText() string {
	return "1/1 spirit"
}

func (s spirit) Description() string {
	return "A 1/1 entity token with OnDeath: Insight 1."
}

func (s spirit) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}

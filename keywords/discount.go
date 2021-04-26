package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discount struct {
	Amount cardobject.SimpleIntValue
}

func (d discount) Validate() error {
	return d.ValidateStruct()
}

func (d discount) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d discount) InteractionText() string {
	return "Discount §Amount."
}

func (d discount) Description() string {
	return "Reduce manacost of all cards in your hand."
}

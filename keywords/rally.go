package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type rally struct{}

func (r rally) ValidateType(root jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(r, root)
}

func (r rally) InteractionText() string {
	return "Rally -> X."
}

func (r rally) Description() string {
	return "Set X to the number of friendly Entities on the Field that share a tag with this card."
}

func (r rally) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}

package keywords

import "github.com/DecentralCardGame/cardobject/jsonschema"

type resetX struct{}

func (r resetX) ValidateType(root jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(r, root)
}

func (r resetX) InteractionText() string {
	return "Reset X."
}

func (r resetX) Description() string {
	return "Set X to 0."
}

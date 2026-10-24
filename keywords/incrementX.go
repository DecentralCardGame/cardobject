package keywords

import "github.com/DecentralCardGame/cardobject/jsonschema"

type incrementX struct{}

func (i incrementX) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(i, r)
}

func (i incrementX) InteractionText() string {
	return "Increment X."
}

func (i incrementX) Description() string {
	return "Increase X by 1."
}

package cardobject

import "github.com/DecentralCardGame/jsonschema"

type actionSelector struct {
	*jsonschema.BasicStruct
	PlayerMode       playerMode
	PlayerCondition  *playerCondition `json:",omitempty"`
	CardMode         cardMode
	ActionConditions *actionConditions `json:",omitempty"`
	Zone             actionZone
	ActionExtractors *actionExtractors `json:",omitempty"`
	AmountExtractor  *intExtractor     `json:",omitempty"`
}

func (a actionSelector) GetInteractionText() string {
	return ""
}

type entitySelector struct {
	*jsonschema.BasicStruct
	PlayerMode       playerMode
	PlayerCondition  *playerCondition `json:",omitempty"`
	CardMode         cardMode
	EntityConditions *entityConditions `json:",omitempty"`
	Zone             entityZone
	EntityExtractors *entityExtractors `json:",omitempty"`
	AmountExtractor  *intExtractor     `json:",omitempty"`
}

func (e entitySelector) GetInteractionText() string {
	return ""
}

type placeSelector struct {
	*jsonschema.BasicStruct
	PlayerMode      playerMode
	PlayerCondition *playerCondition `json:",omitempty"`
	CardMode        cardMode
	PlaceConditions *placeConditions `json:",omitempty"`
	Zone            placeZone
	PlaceExtractors *placeExtractors `json:",omitempty"`
	AmountExtractor *intExtractor    `json:",omitempty"`
}

func (p placeSelector) GetInteractionText() string {
	return ""
}

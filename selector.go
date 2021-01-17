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
	return "Choose §CardMode action §ActionConditions in the §Zone of §PlayerMode player §PlayerCondition. §ActionExtractors §AmountExtractor"
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
	return "Choose §CardMode entity §EntityConditions in the §Zone of §PlayerMode player §PlayerCondition. §EntityExtractors §AmountExtractor"
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
	return "Choose §CardMode place §PlaceConditions in the §Zone of §PlayerMode player §PlayerCondition. §PlaceExtractors §AmountExtractor"
}

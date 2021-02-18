package cardobject

import "github.com/DecentralCardGame/jsonschema"

type actionSelector struct {
	PlayerMode       playerMode
	PlayerCondition  *playerCondition `json:",omitempty"`
	CardMode         cardMode
	ActionConditions *actionConditions `json:",omitempty"`
	Zone             actionZone
	ActionExtractors *actionExtractors `json:",omitempty"`
	AmountExtractor  *intExtractor     `json:",omitempty"`
}

func (a actionSelector) Validate() error {
	return a.ValidateStruct()
}

func (a actionSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionSelector) GetInteractionText() string {
	return "Choose §CardMode action §ActionConditions in the §Zone of §PlayerMode player §PlayerCondition. §ActionExtractors §AmountExtractor"
}

type entitySelector struct {
	PlayerMode       playerMode
	PlayerCondition  *playerCondition `json:",omitempty"`
	CardMode         cardMode
	EntityConditions *entityConditions `json:",omitempty"`
	Zone             entityZone
	EntityExtractors *entityExtractors `json:",omitempty"`
	AmountExtractor  *intExtractor     `json:",omitempty"`
}

func (e entitySelector) Validate() error {
	return e.ValidateStruct()
}

func (e entitySelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entitySelector) GetInteractionText() string {
	return "Choose §CardMode entity §EntityConditions in the §Zone of §PlayerMode player §PlayerCondition. §EntityExtractors §AmountExtractor"
}

type placeSelector struct {
	PlayerMode      playerMode
	PlayerCondition *playerCondition `json:",omitempty"`
	CardMode        cardMode
	PlaceConditions *placeConditions `json:",omitempty"`
	Zone            placeZone
	PlaceExtractors *placeExtractors `json:",omitempty"`
	AmountExtractor *intExtractor    `json:",omitempty"`
}

func (p placeSelector) Validate() error {
	return p.ValidateStruct()
}

func (p placeSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeSelector) GetInteractionText() string {
	return "Choose §CardMode place §PlaceConditions in the §Zone of §PlayerMode player §PlayerCondition. §PlaceExtractors §AmountExtractor"
}

package cardobject

import "github.com/DecentralCardGame/jsonschema"

type ActionSelector struct {
	PlayerMode       PlayerMode
	PlayerCondition  *PlayerCondition `json:",omitempty"`
	CardMode         CardMode
	ActionConditions *ActionConditions `json:",omitempty"`
	Zone             ActionZone
	ActionExtractors *ActionExtractors `json:",omitempty"`
	AmountExtractor  *IntExtractor     `json:",omitempty"`
}

func (a ActionSelector) Validate() error {
	return a.ValidateStruct()
}

func (a ActionSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionSelector) GetInteractionText() string {
	return "Choose §CardMode action §ActionConditions in the §Zone of §PlayerMode player §PlayerCondition. §ActionExtractors §AmountExtractor"
}

type EntitySelector struct {
	PlayerMode       PlayerMode
	PlayerCondition  *PlayerCondition `json:",omitempty"`
	CardMode         CardMode
	EntityConditions *EntityConditions `json:",omitempty"`
	Zone             EntityZone
	EntityExtractors *EntityExtractors `json:",omitempty"`
	AmountExtractor  *IntExtractor     `json:",omitempty"`
}

func (e EntitySelector) Validate() error {
	return e.ValidateStruct()
}

func (e EntitySelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntitySelector) GetInteractionText() string {
	return "Choose §CardMode entity §EntityConditions in the §Zone of §PlayerMode player §PlayerCondition. §EntityExtractors §AmountExtractor"
}

type PlaceSelector struct {
	PlayerMode      PlayerMode
	PlayerCondition *PlayerCondition `json:",omitempty"`
	CardMode        CardMode
	PlaceConditions *PlaceConditions `json:",omitempty"`
	Zone            PlaceZone
	PlaceExtractors *PlaceExtractors `json:",omitempty"`
	AmountExtractor *IntExtractor    `json:",omitempty"`
}

func (p PlaceSelector) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceSelector) GetInteractionText() string {
	return "Choose §CardMode place §PlaceConditions in the §Zone of §PlayerMode player §PlayerCondition. §PlaceExtractors §AmountExtractor"
}

package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type ActionSelector struct {
	PlayerMode       PlayerMode
	PlayerCondition  *PlayerCondition `json:",omitempty"`
	CardMode         CardMode
	ActionConditions *ActionConditions `json:",omitempty"`
	ActionZone       ActionZone
	ActionExtractors *ActionExtractors `json:",omitempty"`
	AmountExtractor  *IntExtractor     `json:",omitempty"`
}

func (a ActionSelector) Validate() error {
	return a.ValidateStruct()
}

func (a ActionSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionSelector) InteractionText() string {
	return "Choose §CardMode action §ActionConditions in the §ActionZone of §PlayerMode player §PlayerCondition. §ActionExtractors §AmountExtractor"
}

type EntitySelector struct {
	PlayerMode       PlayerMode
	PlayerCondition  *PlayerCondition `json:",omitempty"`
	CardMode         CardMode
	EntityConditions *EntityConditions `json:",omitempty"`
	EntityZone       EntityZone
	EntityExtractors *EntityExtractors `json:",omitempty"`
	AmountExtractor  *IntExtractor     `json:",omitempty"`
}

func (e EntitySelector) Validate() error {
	return e.ValidateStruct()
}

func (e EntitySelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntitySelector) InteractionText() string {
	return "Choose §CardMode entity §EntityConditions in the §EntityZone of §PlayerMode player §PlayerCondition. §EntityExtractors §AmountExtractor"
}

type HeadquarterSelector struct {
	PlayerMode            PlayerMode
	PlayerCondition       *PlayerCondition `json:",omitempty"`
	CardMode              CardMode
	HeadquarterConditions *HeadquarterConditions `json:",omitempty"`
	HeadquarterExtractors *HeadquarterExtractors `json:",omitempty"`
	AmountExtractor       *IntExtractor          `json:",omitempty"`
}

func (h HeadquarterSelector) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterSelector) InteractionText() string {
	return "Choose §CardMode action §HeadquarterConditions in the §HeadquarterZone of §PlayerMode player §PlayerCondition. §HeadquarterExtractors §AmountExtractor"
}

type PlaceSelector struct {
	PlayerMode      PlayerMode
	PlayerCondition *PlayerCondition `json:",omitempty"`
	CardMode        CardMode
	PlaceConditions *PlaceConditions `json:",omitempty"`
	PlaceZone       PlaceZone
	PlaceExtractors *PlaceExtractors `json:",omitempty"`
	AmountExtractor *IntExtractor    `json:",omitempty"`
}

func (p PlaceSelector) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceSelector) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceSelector) InteractionText() string {
	return "Choose §CardMode place §PlaceConditions in the §PlaceZone of §PlayerMode player §PlayerCondition. §PlaceExtractors §AmountExtractor"
}

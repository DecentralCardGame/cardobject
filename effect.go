package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type effects []effect

func (e effects) Validate() error {
	return e.ValidateArray()
}

func (e effects) ValidateArray() error {
	return nil
}

func (e effects) GetMinMaxItems() (int, int) {
	return 0, maxAbilityEffectCount
}

type effect struct {
	*jsonschema.BasicInterface
	ProductionEffect *productionEffect `json:",omitempty"`
	DrawEffect       *drawEffect       `json:",omitempty"`
	TokenEffect      *tokenEffect      `json:",omitempty"`
	TargetEffect     *targetEffect     `json:",omitempty"`
	ChooseFromEffect *chooseFromEffect `json:",omitempty"`
}

type productionEffect struct {
	*jsonschema.BasicStruct
	Amount intValue
}

func (p productionEffect) GetInteractionText() string {
	return ""
}

type drawEffect struct {
	*jsonschema.BasicStruct
	DrawAmount intValue
}

func (d drawEffect) GetInteractionText() string {
	return ""
}

type tokenEffect struct {
	*jsonschema.BasicStruct
	TokenAmount intValue
	Token       token
}

func (t tokenEffect) GetInteractionText() string {
	return ""
}

type chooseFromEffect struct {
	*jsonschema.BasicStruct
	Effects effects
}

func (c chooseFromEffect) GetInteractionText() string {
	return ""
}

type targetEffect struct {
	*jsonschema.BasicInterface
	ActionTargetEffect    *actionTargetEffect    `json:",omitempty"`
	EntityTargetEffect    *entityTargetEffect    `json:",omitempty"`
	PlaceTargetEffect     *placeTargetEffect     `json:",omitempty"`
	ExtractorTargetEffect *extractorTargetEffect `json:",omitempty"`
}

type actionTargetEffect struct {
	jsonschema.BasicStruct
	ActionSelector      *actionSelector
	ActionManipulations *actionManipulations
}

func (a actionTargetEffect) GetInteractionText() string {
	return ""
}

type entityTargetEffect struct {
	*jsonschema.BasicStruct
	EntitySelector      *entitySelector
	EntityManipulations *entityManipulations
}

func (e entityTargetEffect) GetInteractionText() string {
	return ""
}

type placeTargetEffect struct {
	*jsonschema.BasicStruct
	PlaceSelector      *placeSelector
	PlaceManipulations *placeManipulations
}

func (p placeTargetEffect) GetInteractionText() string {
	return ""
}

type extractorTargetEffect struct {
	*jsonschema.BasicStruct
	TargetVariable targetVariableName
	Manipulations  *manipulations
}

func (e extractorTargetEffect) GetInteractionText() string {
	return ""
}

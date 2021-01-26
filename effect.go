package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type effects []effect

func (e effects) Validate() error {
	return e.ValidateArray()
}

func (e effects) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
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
	return "Produce §Amount."
}

type drawEffect struct {
	*jsonschema.BasicStruct
	DrawAmount intValue
}

func (d drawEffect) GetInteractionText() string {
	return "Draw §DrawAmount cards."
}

type tokenEffect struct {
	*jsonschema.BasicStruct
	TokenAmount intValue
	Token       token
}

func (t tokenEffect) GetInteractionText() string {
	return "Create §TokenAmount §Token."
}

type chooseFromEffect struct {
	*jsonschema.BasicStruct
	Effects effects
}

func (c chooseFromEffect) GetInteractionText() string {
	return "Do one of §Effects."
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
	return "§ActionSelector §ActionManipulations"
}

type entityTargetEffect struct {
	*jsonschema.BasicStruct
	EntitySelector      *entitySelector
	EntityManipulations *entityManipulations
}

func (e entityTargetEffect) GetInteractionText() string {
	return "§EntitySelector §EntityManipulations"
}

type placeTargetEffect struct {
	*jsonschema.BasicStruct
	PlaceSelector      *placeSelector
	PlaceManipulations *placeManipulations
}

func (p placeTargetEffect) GetInteractionText() string {
	return "§PlaceSelector §PlaceManipulations"
}

type extractorTargetEffect struct {
	*jsonschema.BasicStruct
	TargetVariable targetVariableName
	Manipulations  *manipulations
}

func (e extractorTargetEffect) GetInteractionText() string {
	return "Choose §TargetVariable. §Manipulations"
}

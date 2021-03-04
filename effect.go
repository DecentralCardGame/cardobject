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

func (e effects) GetItemName() string {
	return jsonschema.GetItemNameFromArray(e)
}

type effect struct {
	ProductionEffect *productionEffect `json:",omitempty"`
	WisdomEffect     *wisdomEffect     `json:",omitempty"`
	TokenEffect      *tokenEffect      `json:",omitempty"`
	TargetEffect     *targetEffect     `json:",omitempty"`
	ChooseFromEffect *chooseFromEffect `json:",omitempty"`
}

func (e effect) Validate() error {
	return e.ValidateInterface()
}

func (e effect) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type productionEffect struct {
	Amount intValue
}

func (p productionEffect) Validate() error {
	return p.ValidateStruct()
}

func (p productionEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p productionEffect) GetInteractionText() string {
	return "Produce §Amount."
}

type wisdomEffect struct {
	wisdomAmount intValue
}

func (d wisdomEffect) Validate() error {
	return d.ValidateStruct()
}

func (d wisdomEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d wisdomEffect) GetInteractionText() string {
	return "Gain §WisdomAmount wisdom."
}

type tokenEffect struct {
	TokenAmount intValue
	Token       token
}

func (t tokenEffect) Validate() error {
	return t.ValidateStruct()
}

func (t tokenEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t tokenEffect) GetInteractionText() string {
	return "Create §TokenAmount §Token."
}

type chooseFromEffect struct {
	Effects effects
}

func (c chooseFromEffect) Validate() error {
	return c.ValidateStruct()
}

func (c chooseFromEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c chooseFromEffect) GetInteractionText() string {
	return "Do one of §Effects."
}

type targetEffect struct {
	ActionTargetEffect    *actionTargetEffect    `json:",omitempty"`
	EntityTargetEffect    *entityTargetEffect    `json:",omitempty"`
	PlaceTargetEffect     *placeTargetEffect     `json:",omitempty"`
	ExtractorTargetEffect *extractorTargetEffect `json:",omitempty"`
}

func (t targetEffect) Validate() error {
	return t.ValidateInterface()
}

func (t targetEffect) ValidateInterface() error {
	return jsonschema.ValidateInterface(t)
}

type actionTargetEffect struct {
	ActionSelector      *actionSelector
	ActionManipulations *actionManipulations
}

func (a actionTargetEffect) Validate() error {
	return a.ValidateStruct()
}

func (a actionTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionTargetEffect) GetInteractionText() string {
	return "§ActionSelector §ActionManipulations"
}

type entityTargetEffect struct {
	EntitySelector      *entitySelector
	EntityManipulations *entityManipulations
}

func (e entityTargetEffect) Validate() error {
	return e.ValidateStruct()
}

func (e entityTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityTargetEffect) GetInteractionText() string {
	return "§EntitySelector §EntityManipulations"
}

type placeTargetEffect struct {
	PlaceSelector      *placeSelector
	PlaceManipulations *placeManipulations
}

func (p placeTargetEffect) Validate() error {
	return p.ValidateStruct()
}

func (p placeTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeTargetEffect) GetInteractionText() string {
	return "§PlaceSelector §PlaceManipulations"
}

type extractorTargetEffect struct {
	TargetVariable targetVariableName
	Manipulations  *manipulations
}

func (e extractorTargetEffect) Validate() error {
	return e.ValidateStruct()
}

func (e extractorTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e extractorTargetEffect) GetInteractionText() string {
	return "Choose §TargetVariable. §Manipulations"
}

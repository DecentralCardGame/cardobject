package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type Effects []Effect

func (e Effects) Validate(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e Effects) MinMaxItems() (int, int) {
	return 0, maxAbilityEffectCount
}

func (e Effects) ItemName() string {
	return jsonschema.ItemNameFromArray(e)
}

type Effect struct {
	GrowthEffect     *GrowthEffect     `json:",omitempty"`
	ProductionEffect *ProductionEffect `json:",omitempty"`
	WisdomEffect     *WisdomEffect     `json:",omitempty"`
	TokenEffect      *TokenEffect      `json:",omitempty"`
	TargetEffect     *TargetEffect     `json:",omitempty"`
	ChooseFromEffect *ChooseFromEffect `json:",omitempty"`
}

func (e Effect) Validate(r jsonschema.RootElement) error {
	implementer, err := e.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.Validate(r)
}

func (e Effect) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(e)
}

type GrowthEffect struct {
	Keyword      *Keyword `json:",omitempty"`
	GrowthAmount IntValue
}

func (g GrowthEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(g, r)
}

func (g GrowthEffect) InteractionText() string {
	return "Gain §GrowthAmount growth. (§Keyword)"
}

type ProductionEffect struct {
	Keyword *Keyword `json:",omitempty"`
	Amount  IntValue
}

func (p ProductionEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p ProductionEffect) InteractionText() string {
	return "Produce §Amount. (§Keyword)"
}

type WisdomEffect struct {
	Keyword      *Keyword `json:",omitempty"`
	wisdomAmount IntValue
}

func (d WisdomEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d WisdomEffect) InteractionText() string {
	return "Gain §WisdomAmount wisdom. (§Keyword)"
}

type TokenEffect struct {
	Keyword     *Keyword `json:",omitempty"`
	TokenAmount IntValue
	Token       Token
}

func (t TokenEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t TokenEffect) InteractionText() string {
	return "Create §TokenAmount §Token. (§Keyword)"
}

type ChooseFromEffect struct {
	Effects Effects
}

func (c ChooseFromEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c ChooseFromEffect) InteractionText() string {
	return "Do one of §Effects."
}

type TargetEffect struct {
	ActionTargetEffect      *ActionTargetEffect      `json:",omitempty"`
	EntityTargetEffect      *EntityTargetEffect      `json:",omitempty"`
	HeadquarterTargetEffect *HeadquarterTargetEffect `json:",omitempty"`
	PlaceTargetEffect       *PlaceTargetEffect       `json:",omitempty"`
	ExtractorTargetEffect   *ExtractorTargetEffect   `json:",omitempty"`
}

func (t TargetEffect) Validate(r jsonschema.RootElement) error {
	implementer, err := t.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.Validate(r)
}

func (t TargetEffect) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(t)
}

type ActionTargetEffect struct {
	ActionSelector      *ActionSelector
	ActionManipulations *ActionManipulations
}

func (a ActionTargetEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionTargetEffect) InteractionText() string {
	return "§ActionSelector §ActionManipulations"
}

type EntityTargetEffect struct {
	EntitySelector      *EntitySelector
	EntityManipulations *EntityManipulations
}

func (e EntityTargetEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityTargetEffect) InteractionText() string {
	return "§EntitySelector §EntityManipulations"
}

type HeadquarterTargetEffect struct {
	HeadquarterSelector      *HeadquarterSelector
	HeadquarterManipulations *HeadquarterManipulations
}

func (h HeadquarterTargetEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterTargetEffect) InteractionText() string {
	return "§HeadquarterSelector §HeadquarterManipulations"
}

type PlaceTargetEffect struct {
	PlaceSelector      *PlaceSelector
	PlaceManipulations *PlaceManipulations
}

func (p PlaceTargetEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceTargetEffect) InteractionText() string {
	return "§PlaceSelector §PlaceManipulations"
}

type ExtractorTargetEffect struct {
	TargetVariable TargetVariableName
	Manipulations  *Manipulations
}

func (e ExtractorTargetEffect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e ExtractorTargetEffect) InteractionText() string {
	return "Choose §TargetVariable. §Manipulations"
}

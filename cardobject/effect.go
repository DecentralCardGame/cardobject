package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type Effects []Effect

func (e Effects) Validate() error {
	return e.ValidateArray()
}

func (e Effects) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
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

func (e Effect) Validate() error {
	return e.ValidateInterface()
}

func (e Effect) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type GrowthEffect struct {
	Keyword      *Keyword `json:",omitempty"`
	GrowthAmount IntValue
}

func (g GrowthEffect) Validate() error {
	return g.ValidateStruct()
}

func (g GrowthEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(g)
}

func (g GrowthEffect) InteractionText() string {
	return "Gain §GrowthAmount growth. (§Keyword)"
}

type ProductionEffect struct {
	Keyword *Keyword `json:",omitempty"`
	Amount  IntValue
}

func (p ProductionEffect) Validate() error {
	return p.ValidateStruct()
}

func (p ProductionEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p ProductionEffect) InteractionText() string {
	return "Produce §Amount. (§Keyword)"
}

type WisdomEffect struct {
	Keyword      *Keyword `json:",omitempty"`
	wisdomAmount IntValue
}

func (d WisdomEffect) Validate() error {
	return d.ValidateStruct()
}

func (d WisdomEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d WisdomEffect) InteractionText() string {
	return "Gain §WisdomAmount wisdom. (§Keyword)"
}

type TokenEffect struct {
	Keyword     *Keyword `json:",omitempty"`
	TokenAmount IntValue
	Token       Token
}

func (t TokenEffect) Validate() error {
	return t.ValidateStruct()
}

func (t TokenEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t TokenEffect) InteractionText() string {
	return "Create §TokenAmount §Token. (§Keyword)"
}

type ChooseFromEffect struct {
	Effects Effects
}

func (c ChooseFromEffect) Validate() error {
	return c.ValidateStruct()
}

func (c ChooseFromEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
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

func (t TargetEffect) Validate() error {
	return t.ValidateInterface()
}

func (t TargetEffect) ValidateInterface() error {
	return jsonschema.ValidateInterface(t)
}

type ActionTargetEffect struct {
	ActionSelector      *ActionSelector
	ActionManipulations *ActionManipulations
}

func (a ActionTargetEffect) Validate() error {
	return a.ValidateStruct()
}

func (a ActionTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionTargetEffect) InteractionText() string {
	return "§ActionSelector §ActionManipulations"
}

type EntityTargetEffect struct {
	EntitySelector      *EntitySelector
	EntityManipulations *EntityManipulations
}

func (e EntityTargetEffect) Validate() error {
	return e.ValidateStruct()
}

func (e EntityTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityTargetEffect) InteractionText() string {
	return "§EntitySelector §EntityManipulations"
}

type HeadquarterTargetEffect struct {
	HeadquarterSelector      *HeadquarterSelector
	HeadquarterManipulations *HeadquarterManipulations
}

func (h HeadquarterTargetEffect) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterTargetEffect) InteractionText() string {
	return "§HeadquarterSelector §HeadquarterManipulations"
}

type PlaceTargetEffect struct {
	PlaceSelector      *PlaceSelector
	PlaceManipulations *PlaceManipulations
}

func (p PlaceTargetEffect) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceTargetEffect) InteractionText() string {
	return "§PlaceSelector §PlaceManipulations"
}

type ExtractorTargetEffect struct {
	TargetVariable TargetVariableName
	Manipulations  *Manipulations
}

func (e ExtractorTargetEffect) Validate() error {
	return e.ValidateStruct()
}

func (e ExtractorTargetEffect) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e ExtractorTargetEffect) InteractionText() string {
	return "Choose §TargetVariable. §Manipulations"
}

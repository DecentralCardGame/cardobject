package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type Manipulations struct {
	ActionManipulations      *ActionManipulations      `json:",omitempty"`
	EntityManipulations      *EntityManipulations      `json:",omitempty"`
	HeadquarterManipulations *HeadquarterManipulations `json:",omitempty"`
	PlaceManipulations       *PlaceManipulations       `json:",omitempty"`
}

func (m Manipulations) ValidateType(r jsonschema.RootElement) error {
	implementer, err := m.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (m Manipulations) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(m)
}

type ActionManipulations []ActionManipulation

func (a ActionManipulations) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range a {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a ActionManipulations) MinMaxItems() (int, int) {
	return 0, 3
}

func (a ActionManipulations) ItemName() string {
	return jsonschema.ItemNameFromArray(a)
}

type EntityManipulations []EntityManipulation

func (e EntityManipulations) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range e {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e EntityManipulations) MinMaxItems() (int, int) {
	return 0, 3
}

func (e EntityManipulations) ItemName() string {
	return jsonschema.ItemNameFromArray(e)
}

type HeadquarterManipulations []HeadquarterManipulation

func (h HeadquarterManipulations) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range h {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (h HeadquarterManipulations) MinMaxItems() (int, int) {
	return 0, 3
}

func (h HeadquarterManipulations) ItemName() string {
	return jsonschema.ItemNameFromArray(h)
}

type PlaceManipulations []PlaceManipulation

func (p PlaceManipulations) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range p {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p PlaceManipulations) MinMaxItems() (int, int) {
	return 0, 3
}

func (p PlaceManipulations) ItemName() string {
	return jsonschema.ItemNameFromArray(p)
}

type ActionManipulation struct {
	ActionEffectManipulation *ActionEffectManipulation `json:",omitempty"`
	ActionIntManipulation    *ActionIntManipulation    `json:",omitempty"`
	ActionStringManipulation *ActionStringManipulation `json:",omitempty"`
	ActionTagManipulation    *ActionTagManipulation    `json:",omitempty"`
	ActionZoneChange         *ActionZoneChange         `json:",omitempty"`
}

func (a ActionManipulation) ValidateType(r jsonschema.RootElement) error {
	implementer, err := a.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (a ActionManipulation) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(a)
}

type EntityManipulation struct {
	EntityAbilityManipulation *EntityAbilityManipulation `json:",omitempty"`
	EntityIntManipulation     *EntityIntManipulation     `json:",omitempty"`
	EntityStringManipulation  *EntityStringManipulation  `json:",omitempty"`
	EntityTagManipulation     *EntityTagManipulation     `json:",omitempty"`
	EntityZoneChange          *EntityZoneChange          `json:",omitempty"`
}

func (e EntityManipulation) ValidateType(r jsonschema.RootElement) error {
	implementer, err := e.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (e EntityManipulation) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(e)
}

type HeadquarterManipulation struct {
	HeadquarterEffectManipulation *HeadquarterAbilityManipulation `json:",omitempty"`
	HeadquarterIntManipulation    *HeadquarterIntManipulation     `json:",omitempty"`
	HeadquarterStringManipulation *HeadquarterStringManipulation  `json:",omitempty"`
	HeadquarterTagManipulation    *HeadquarterTagManipulation     `json:",omitempty"`
}

func (h HeadquarterManipulation) ValidateType(r jsonschema.RootElement) error {
	implementer, err := h.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (h HeadquarterManipulation) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(h)
}

type PlaceManipulation struct {
	PlaceAbilityManipulation *PlaceAbilityManipulation `json:",omitempty"`
	PlaceIntManipulation     *PlaceIntManipulation     `json:",omitempty"`
	PlaceStringManipulation  *PlaceStringManipulation  `json:",omitempty"`
	PlaceTagManipulation     *PlaceTagManipulation     `json:",omitempty"`
	PlaceZoneChange          *PlaceZoneChange          `json:",omitempty"`
}

func (p PlaceManipulation) ValidateType(r jsonschema.RootElement) error {
	implementer, err := p.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (p PlaceManipulation) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(p)
}

type ActionEffectManipulation struct {
	Effect         *Effect `json:",omitempty"`
	EffectOperator AbilityEffectOperator
	Keyword        *Keyword `json:",omitempty"`
}

func (a ActionEffectManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionEffectManipulation) InteractionText() string {
	return "It §EffectOperator §Effect.(§Keyword)"
}

type ActionIntManipulation struct {
	IntProperty ActionIntProperty
	IntOperator IntOperator
	IntValue    IntValue
	Keyword     *Keyword `json:",omitempty"`
}

func (a ActionIntManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionIntManipulation) InteractionText() string {
	return "§IntOperator §IntProperty §IntValue.(§Keyword)"
}

type ActionStringManipulation struct {
	StringProperty ActionStringProperty
	StringOperator StringOperator
	StringValue    SimpleStringValue
	Keyword        *Keyword `json:",omitempty"`
}

func (a ActionStringManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type ActionTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (a ActionTagManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionTagManipulation) InteractionText() string {
	return "§TagOperator tag §TagValue.(§Keyword)"
}

type EntityAbilityManipulation struct {
	Ability         *Ability `json:",omitempty"`
	AbilityOperator AbilityEffectOperator
	Keyword         *Keyword `json:",omitempty"`
}

func (e EntityAbilityManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityAbilityManipulation) InteractionText() string {
	return "It §AbilityOperator §Ability.(§Keyword)"
}

type EntityIntManipulation struct {
	IntProperty EntityIntProperty
	IntOperator IntOperator
	IntValue    IntValue
	Keyword     *Keyword `json:",omitempty"`
}

func (e EntityIntManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityIntManipulation) InteractionText() string {
	return "§IntOperator §IntProperty §IntValue.(§Keyword)"
}

type EntityStringManipulation struct {
	StringProperty EntityStringProperty
	StringOperator StringOperator
	StringValue    SimpleStringValue
	Keyword        *Keyword `json:",omitempty"`
}

func (e EntityStringManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type EntityTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (e EntityTagManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityTagManipulation) InteractionText() string {
	return "It §TagOperator §TagValue.(§Keyword)"
}

type HeadquarterAbilityManipulation struct {
	Ability         *Ability `json:",omitempty"`
	AbilityOperator AbilityEffectOperator
	Keyword         *Keyword `json:",omitempty"`
}

func (h HeadquarterAbilityManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterAbilityManipulation) InteractionText() string {
	return "It §AbilityOperator §Effect.(§Keyword)"
}

type HeadquarterIntManipulation struct {
	IntProperty HeadquarterIntProperty
	IntOperator IntOperator
	IntValue    IntValue
	Keyword     *Keyword `json:",omitempty"`
}

func (h HeadquarterIntManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterIntManipulation) InteractionText() string {
	return "§IntOperator §IntProperty §IntValue.(§Keyword)"
}

type HeadquarterStringManipulation struct {
	StringProperty HeadquarterStringProperty
	StringOperator StringOperator
	StringValue    SimpleStringValue
	Keyword        *Keyword `json:",omitempty"`
}

func (h HeadquarterStringManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type HeadquarterTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (h HeadquarterTagManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterTagManipulation) InteractionText() string {
	return "§TagOperator tag §TagValue.(§Keyword)"
}

type PlaceAbilityManipulation struct {
	Ability         *Ability `json:",omitempty"`
	AbilityOperator AbilityEffectOperator
	Keyword         *Keyword `json:",omitempty"`
}

func (p PlaceAbilityManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceAbilityManipulation) InteractionText() string {
	return "It §AbilityOperator §Ability.(§Keyword)"
}

type PlaceIntManipulation struct {
	IntProperty PlaceIntProperty
	IntOperator IntOperator
	IntValue    IntValue
	Keyword     *Keyword `json:",omitempty"`
}

func (p PlaceIntManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceIntManipulation) InteractionText() string {
	return "§IntOperator §IntProperty §IntValue.(§Keyword)"
}

type PlaceStringManipulation struct {
	StringProperty PlaceStringProperty
	StringOperator StringOperator
	StringValue    SimpleStringValue
	Keyword        *Keyword `json:",omitempty"`
}

func (p PlaceStringManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type PlaceTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (p PlaceTagManipulation) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceTagManipulation) InteractionText() string {
	return "It §TagOperator §TagValue.(§Keyword)"
}

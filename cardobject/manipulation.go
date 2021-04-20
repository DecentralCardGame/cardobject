package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type Manipulations struct {
	ActionManipulations      *ActionManipulations      `json:",omitempty"`
	EntityManipulations      *EntityManipulations      `json:",omitempty"`
	HeadquarterManipulations *HeadquarterManipulations `json:",omitempty"`
	PlaceManipulations       *PlaceManipulations       `json:",omitempty"`
}

func (m Manipulations) Validate() error {
	return m.ValidateInterface()
}

func (m Manipulations) ValidateInterface() error {
	return jsonschema.ValidateInterface(m)
}

type ActionManipulations []ActionManipulation

func (a ActionManipulations) Validate() error {
	return a.ValidateArray()
}

func (a ActionManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
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

func (e EntityManipulations) Validate() error {
	return e.ValidateArray()
}

func (e EntityManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
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

func (h HeadquarterManipulations) Validate() error {
	return h.ValidateArray()
}

func (h HeadquarterManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range h {
		err := v.Validate()
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

func (p PlaceManipulations) Validate() error {
	return p.ValidateArray()
}

func (p PlaceManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
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

func (a ActionManipulation) Validate() error {
	return a.ValidateInterface()
}

func (a ActionManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type EntityManipulation struct {
	EntityAbilityManipulation *EntityAbilityManipulation `json:",omitempty"`
	EntityIntManipulation     *EntityIntManipulation     `json:",omitempty"`
	EntityStringManipulation  *EntityStringManipulation  `json:",omitempty"`
	EntityTagManipulation     *EntityTagManipulation     `json:",omitempty"`
	EntityZoneChange          *EntityZoneChange          `json:",omitempty"`
}

func (e EntityManipulation) Validate() error {
	return e.ValidateInterface()
}

func (e EntityManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type HeadquarterManipulation struct {
	HeadquarterEffectManipulation *HeadquarterAbilityManipulation `json:",omitempty"`
	HeadquarterIntManipulation    *HeadquarterIntManipulation     `json:",omitempty"`
	HeadquarterStringManipulation *HeadquarterStringManipulation  `json:",omitempty"`
	HeadquarterTagManipulation    *HeadquarterTagManipulation     `json:",omitempty"`
}

func (h HeadquarterManipulation) Validate() error {
	return h.ValidateInterface()
}

func (h HeadquarterManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(h)
}

type PlaceManipulation struct {
	PlaceAbilityManipulation *PlaceAbilityManipulation `json:",omitempty"`
	PlaceIntManipulation     *PlaceIntManipulation     `json:",omitempty"`
	PlaceStringManipulation  *PlaceStringManipulation  `json:",omitempty"`
	PlaceTagManipulation     *PlaceTagManipulation     `json:",omitempty"`
	PlaceZoneChange          *PlaceZoneChange          `json:",omitempty"`
}

func (p PlaceManipulation) Validate() error {
	return p.ValidateInterface()
}

func (p PlaceManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type ActionEffectManipulation struct {
	Effect         *Effect `json:",omitempty"`
	EffectOperator AbilityEffectOperator
	Keyword        *Keyword `json:",omitempty"`
}

func (a ActionEffectManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a ActionEffectManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
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

func (a ActionIntManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a ActionIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
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

func (a ActionStringManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a ActionStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type ActionTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (a ActionTagManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a ActionTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionTagManipulation) InteractionText() string {
	return "§TagOperator tag §TagValue.(§Keyword)"
}

type EntityAbilityManipulation struct {
	Ability         *Ability `json:",omitempty"`
	AbilityOperator AbilityEffectOperator
	Keyword         *Keyword `json:",omitempty"`
}

func (e EntityAbilityManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e EntityAbilityManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
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

func (e EntityIntManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e EntityIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
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

func (e EntityStringManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e EntityStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type EntityTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (e EntityTagManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e EntityTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityTagManipulation) InteractionText() string {
	return "It §TagOperator §TagValue.(§Keyword)"
}

type HeadquarterAbilityManipulation struct {
	Ability         *Ability `json:",omitempty"`
	AbilityOperator AbilityEffectOperator
	Keyword         *Keyword `json:",omitempty"`
}

func (h HeadquarterAbilityManipulation) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterAbilityManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
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

func (h HeadquarterIntManipulation) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
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

func (h HeadquarterStringManipulation) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type HeadquarterTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (h HeadquarterTagManipulation) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterTagManipulation) InteractionText() string {
	return "§TagOperator tag §TagValue.(§Keyword)"
}

type PlaceAbilityManipulation struct {
	Ability         *Ability `json:",omitempty"`
	AbilityOperator AbilityEffectOperator
	Keyword         *Keyword `json:",omitempty"`
}

func (p PlaceAbilityManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceAbilityManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
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

func (p PlaceIntManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
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

func (p PlaceStringManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceStringManipulation) InteractionText() string {
	return "§StringOperator §StringProperty §StringValue.(§Keyword)"
}

type PlaceTagManipulation struct {
	TagValue    Tag
	TagOperator StringOperator
	Keyword     *Keyword `json:",omitempty"`
}

func (p PlaceTagManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceTagManipulation) InteractionText() string {
	return "It §TagOperator §TagValue.(§Keyword)"
}

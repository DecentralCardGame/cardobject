package cardobject

import "github.com/DecentralCardGame/jsonschema"

type manipulations struct {
	*jsonschema.BasicInterface
	ActionManipulations *actionManipulations `json:",omitempty"`
	EntityManipulations *entityManipulations `json:",omitempty"`
	PlaceManipulations  *placeManipulations  `json:",omitempty"`
}

type actionManipulations []actionManipulation

func (a actionManipulations) Validate() error {
	return a.ValidateArray()
}

func (a actionManipulations) ValidateArray() error {
	return nil
}

func (a actionManipulations) GetMinMaxItems() (int, int) {
	return 0, 3
}

type entityManipulations []entityManipulation

func (e entityManipulations) Validate() error {
	return e.ValidateArray()
}

func (e entityManipulations) ValidateArray() error {
	return nil
}

func (e entityManipulations) GetMinMaxItems() (int, int) {
	return 0, 3
}

type placeManipulations []placeManipulation

func (p placeManipulations) Validate() error {
	return p.ValidateArray()
}

func (p placeManipulations) ValidateArray() error {
	return nil
}

func (p placeManipulations) GetMinMaxItems() (int, int) {
	return 0, 3
}

type actionManipulation struct {
	*jsonschema.BasicInterface
	ActionEffectManipulation *actionEffectManipulation `json:",omitempty"`
	ActionIntManipulation    *actionIntManipulation    `json:",omitempty"`
	ActionStringManipulation *actionStringManipulation `json:",omitempty"`
	ActionTagManipulation    *actionTagManipulation    `json:",omitempty"`
	ActionZoneChange         *actionZoneChange         `json:",omitempty"`
}

type entityManipulation struct {
	*jsonschema.BasicInterface
	EntityAbilityManipulation *entityAbilityManipulation `json:",omitempty"`
	EntityIntManipulation     *entityIntManipulation     `json:",omitempty"`
	EntityStringManipulation  *entityStringManipulation  `json:",omitempty"`
	EntityTagManipulation     *entityTagManipulation     `json:",omitempty"`
	EntityZoneChange          *entityZoneChange          `json:",omitempty"`
}

type placeManipulation struct {
	*jsonschema.BasicInterface
	PlaceAbilityManipulation *placeAbilityManipulation `json:",omitempty"`
	PlaceIntManipulation     *placeIntManipulation     `json:",omitempty"`
	PlaceStringManipulation  *placeStringManipulation  `json:",omitempty"`
	PlaceTagManipulation     *placeTagManipulation     `json:",omitempty"`
	PlaceZoneChange          *placeZoneChange          `json:",omitempty"`
}

type actionEffectManipulation struct {
	*jsonschema.BasicStruct
	Effect         effect
	EffectOperator abilityEffectOperator
}

func (a actionEffectManipulation) GetInteractionText() string {
	return ""
}

type actionIntManipulation struct {
	*jsonschema.BasicStruct
	IntProperty actionIntProperty
	IntOperator intOperator
	IntValue    intValue
}

func (a actionIntManipulation) GetInteractionText() string {
	return ""
}

type actionStringManipulation struct {
	*jsonschema.BasicStruct
	StringProperty actionStringProperty
	StringOperator stringOperator
	StringValue    simpleStringValue
}

func (a actionStringManipulation) GetInteractionText() string {
	return ""
}

type actionTagManipulation struct {
	*jsonschema.BasicStruct
	StringValue    tag
	StringOperator stringOperator
}

func (a actionTagManipulation) GetInteractionText() string {
	return ""
}

type entityAbilityManipulation struct {
	*jsonschema.BasicStruct
	Ability         ability
	AbilityOperator abilityEffectOperator
}

func (e entityAbilityManipulation) GetInteractionText() string {
	return ""
}

type entityIntManipulation struct {
	*jsonschema.BasicStruct
	IntProperty entityIntProperty
	IntOperator intOperator
	IntValue    intValue
}

func (e entityIntManipulation) GetInteractionText() string {
	return ""
}

type entityStringManipulation struct {
	*jsonschema.BasicStruct
	StringProperty entityStringProperty
	StringOperator stringOperator
	StringValue    simpleStringValue
}

func (e entityStringManipulation) GetInteractionText() string {
	return ""
}

type entityTagManipulation struct {
	*jsonschema.BasicStruct
	StringValue    tag
	StringOperator stringOperator
}

func (e entityTagManipulation) GetInteractionText() string {
	return ""
}

type placeAbilityManipulation struct {
	*jsonschema.BasicStruct
	Ability         ability
	AbilityOperator abilityEffectOperator
}

func (p placeAbilityManipulation) GetInteractionText() string {
	return ""
}

type placeIntManipulation struct {
	*jsonschema.BasicStruct
	IntProperty placeIntProperty
	IntOperator intOperator
	IntValue    intValue
}

func (p placeIntManipulation) GetInteractionText() string {
	return ""
}

type placeStringManipulation struct {
	*jsonschema.BasicStruct
	StringProperty placeStringProperty
	StringOperator stringOperator
	StringValue    simpleStringValue
}

func (p placeStringManipulation) GetInteractionText() string {
	return ""
}

type placeTagManipulation struct {
	*jsonschema.BasicStruct
	StringValue    tag
	StringOperator stringOperator
}

func (p placeTagManipulation) GetInteractionText() string {
	return ""
}

package cardobject

import "github.com/DecentralCardGame/jsonschema"

type manipulations struct {
	ActionManipulations *actionManipulations `json:",omitempty"`
	EntityManipulations *entityManipulations `json:",omitempty"`
	PlaceManipulations  *placeManipulations  `json:",omitempty"`
}

func (m manipulations) Validate() error {
	return m.ValidateInterface()
}

func (m manipulations) ValidateInterface() error {
	return jsonschema.ValidateInterface(m)
}

type actionManipulations []actionManipulation

func (a actionManipulations) Validate() error {
	return a.ValidateArray()
}

func (a actionManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a actionManipulations) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (a actionManipulations) GetItemName() string {
	return jsonschema.GetItemNameFromArray(a)
}

type entityManipulations []entityManipulation

func (e entityManipulations) Validate() error {
	return e.ValidateArray()
}

func (e entityManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e entityManipulations) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (e entityManipulations) GetItemName() string {
	return jsonschema.GetItemNameFromArray(e)
}

type placeManipulations []placeManipulation

func (p placeManipulations) Validate() error {
	return p.ValidateArray()
}

func (p placeManipulations) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p placeManipulations) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (p placeManipulations) GetItemName() string {
	return jsonschema.GetItemNameFromArray(p)
}

type actionManipulation struct {
	ActionEffectManipulation *actionEffectManipulation `json:",omitempty"`
	ActionIntManipulation    *actionIntManipulation    `json:",omitempty"`
	ActionStringManipulation *actionStringManipulation `json:",omitempty"`
	ActionTagManipulation    *actionTagManipulation    `json:",omitempty"`
	ActionZoneChange         *actionZoneChange         `json:",omitempty"`
}

func (a actionManipulation) Validate() error {
	return a.ValidateInterface()
}

func (a actionManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type entityManipulation struct {
	EntityAbilityManipulation *entityAbilityManipulation `json:",omitempty"`
	EntityIntManipulation     *entityIntManipulation     `json:",omitempty"`
	EntityStringManipulation  *entityStringManipulation  `json:",omitempty"`
	EntityTagManipulation     *entityTagManipulation     `json:",omitempty"`
	EntityZoneChange          *entityZoneChange          `json:",omitempty"`
}

func (e entityManipulation) Validate() error {
	return e.ValidateInterface()
}

func (e entityManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type placeManipulation struct {
	PlaceAbilityManipulation *placeAbilityManipulation `json:",omitempty"`
	PlaceIntManipulation     *placeIntManipulation     `json:",omitempty"`
	PlaceStringManipulation  *placeStringManipulation  `json:",omitempty"`
	PlaceTagManipulation     *placeTagManipulation     `json:",omitempty"`
	PlaceZoneChange          *placeZoneChange          `json:",omitempty"`
}

func (p placeManipulation) Validate() error {
	return p.ValidateInterface()
}

func (p placeManipulation) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type actionEffectManipulation struct {
	Effect         effect
	EffectOperator abilityEffectOperator
}

func (a actionEffectManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a actionEffectManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionEffectManipulation) GetInteractionText() string {
	return "It §EffectOperator §Effect."
}

type actionIntManipulation struct {
	IntProperty actionIntProperty
	IntOperator intOperator
	IntValue    intValue
}

func (a actionIntManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a actionIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionIntManipulation) GetInteractionText() string {
	return "§IntOperator §IntProperty §IntValue."
}

type actionStringManipulation struct {
	StringProperty actionStringProperty
	StringOperator stringOperator
	StringValue    simpleStringValue
}

func (a actionStringManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a actionStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionStringManipulation) GetInteractionText() string {
	return "§StringOperator §StringProperty §StringValue."
}

type actionTagManipulation struct {
	TagValue    tag
	TagOperator stringOperator
}

func (a actionTagManipulation) Validate() error {
	return a.ValidateStruct()
}

func (a actionTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionTagManipulation) GetInteractionText() string {
	return "§TagOperator tag §TagValue."
}

type entityAbilityManipulation struct {
	Ability         ability
	AbilityOperator abilityEffectOperator
}

func (e entityAbilityManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e entityAbilityManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityAbilityManipulation) GetInteractionText() string {
	return "It §AbilityOperator §Ability."
}

type entityIntManipulation struct {
	IntProperty entityIntProperty
	IntOperator intOperator
	IntValue    intValue
}

func (e entityIntManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e entityIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityIntManipulation) GetInteractionText() string {
	return "§IntOperator §IntProperty §IntValue"
}

type entityStringManipulation struct {
	StringProperty entityStringProperty
	StringOperator stringOperator
	StringValue    simpleStringValue
}

func (e entityStringManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e entityStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityStringManipulation) GetInteractionText() string {
	return "§StringOperator §StringProperty §StringValue."
}

type entityTagManipulation struct {
	TagValue    tag
	TagOperator stringOperator
}

func (e entityTagManipulation) Validate() error {
	return e.ValidateStruct()
}

func (e entityTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityTagManipulation) GetInteractionText() string {
	return "It §TagOperator §TagValue."
}

type placeAbilityManipulation struct {
	Ability         ability
	AbilityOperator abilityEffectOperator
}

func (p placeAbilityManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p placeAbilityManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeAbilityManipulation) GetInteractionText() string {
	return "It §AbilityOperator §Ability."
}

type placeIntManipulation struct {
	IntProperty placeIntProperty
	IntOperator intOperator
	IntValue    intValue
}

func (p placeIntManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p placeIntManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeIntManipulation) GetInteractionText() string {
	return "§IntOperator §IntProperty §IntValue."
}

type placeStringManipulation struct {
	StringProperty placeStringProperty
	StringOperator stringOperator
	StringValue    simpleStringValue
}

func (p placeStringManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p placeStringManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeStringManipulation) GetInteractionText() string {
	return "§StringOperator §StringProperty §StringValue."
}

type placeTagManipulation struct {
	TagValue    tag
	TagOperator stringOperator
}

func (p placeTagManipulation) Validate() error {
	return p.ValidateStruct()
}

func (p placeTagManipulation) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeTagManipulation) GetInteractionText() string {
	return "It §TagOperator §TagValue."
}

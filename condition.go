package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type CardConditions struct {
	ActionConditions      *ActionConditions      `json:",omitempty"`
	EntityConditions      *EntityConditions      `json:",omitempty"`
	HeadquarterConditions *HeadquarterConditions `json:",omitempty"`
	PlaceConditions       *PlaceConditions       `json:",omitempty"`
	ThisConditions        *ThisCondition         `json:",omitempty"`
}

func (c CardConditions) Validate() error {
	return c.ValidateInterface()
}

func (c CardConditions) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type ActionConditions []ActionCondition

func (a ActionConditions) Validate() error {
	return a.ValidateArray()
}

func (a ActionConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a ActionConditions) ItemName() string {
	return jsonschema.ItemNameFromArray(a)
}

func (a ActionConditions) MinMaxItems() (int, int) {
	return 0, 3
}

type EntityConditions []EntityCondition

func (e EntityConditions) Validate() error {
	return e.ValidateArray()
}

func (e EntityConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e EntityConditions) MinMaxItems() (int, int) {
	return 0, 3
}

func (e EntityConditions) ItemName() string {
	return jsonschema.ItemNameFromArray(e)
}

type HeadquarterConditions []HeadquarterCondition

func (h HeadquarterConditions) Validate() error {
	return h.ValidateArray()
}

func (h HeadquarterConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range h {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (h HeadquarterConditions) MinMaxItems() (int, int) {
	return 0, 3
}

func (h HeadquarterConditions) ItemName() string {
	return jsonschema.ItemNameFromArray(h)
}

type PlaceConditions []PlaceCondition

func (p PlaceConditions) Validate() error {
	return p.ValidateArray()
}

func (p PlaceConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p PlaceConditions) MinMaxItems() (int, int) {
	return 0, 3
}

func (p PlaceConditions) ItemName() string {
	return jsonschema.ItemNameFromArray(p)
}

type PlayerConditions []PlayerCondition

func (p PlayerConditions) Validate() error {
	return p.ValidateArray()
}

func (p PlayerConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p PlayerConditions) MinMaxItems() (int, int) {
	return 0, 3
}

func (p PlayerConditions) ItemName() string {
	return jsonschema.ItemNameFromArray(p)
}

type ActionCondition struct {
	ActionIntCondition    *ActionIntCondition    `json:",omitempty"`
	ActionStringCondition *ActionStringCondition `json:",omitempty"`
	ActionTagCondition    *ActionTagCondition    `json:",omitempty"`
}

func (a ActionCondition) Validate() error {
	return a.ValidateInterface()
}

func (a ActionCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type EntityCondition struct {
	EntityIntCondition    *EntityIntCondition    `json:",omitempty"`
	EntityStringCondition *EntityStringCondition `json:",omitempty"`
	EntityTagCondition    *EntityTagCondition    `json:",omitempty"`
}

func (e EntityCondition) Validate() error {
	return e.ValidateInterface()
}

func (e EntityCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type HeadquarterCondition struct {
	HeadquarterIntCondition    *HeadquarterIntCondition    `json:",omitempty"`
	HeadquarterStringCondition *HeadquarterStringCondition `json:",omitempty"`
	HeadquarterTagCondition    *HeadquarterTagCondition    `json:",omitempty"`
}

func (h HeadquarterCondition) Validate() error {
	return h.ValidateInterface()
}

func (h HeadquarterCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(h)
}

type PlaceCondition struct {
	PlaceIntCondition    *PlaceIntCondition    `json:",omitempty"`
	PlaceStringCondition *PlaceStringCondition `json:",omitempty"`
	PlaceTagCondition    *PlaceTagCondition    `json:",omitempty"`
}

func (p PlaceCondition) Validate() error {
	return p.ValidateInterface()
}

func (p PlaceCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type PlayerCondition struct {
	PlayerIntCondition *PlayerIntCondition `json:",omitempty"`
}

func (p PlayerCondition) Validate() error {
	return p.ValidateInterface()
}

func (p PlayerCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type ThisCondition struct{}

func (t ThisCondition) Validate() error {
	return t.ValidateStruct()
}

func (t ThisCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t ThisCondition) InteractionText() string {
	return "this"
}

type ActionIntCondition struct {
	ActionIntProperty ActionIntProperty
	IntValue          SimpleIntValue
	IntComparator     IntComparator
}

func (a ActionIntCondition) Validate() error {
	return a.ValidateStruct()
}

func (a ActionIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionIntCondition) InteractionText() string {
	return "with §ActionIntProperty §IntComparator §IntValue"
}

type ActionStringCondition struct {
	ActionStringProperty ActionStringProperty
	StringValue          SimpleStringValue
	StringComparator     StringComparator
}

func (a ActionStringCondition) Validate() error {
	return a.ValidateStruct()
}

func (a ActionStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionStringCondition) InteractionText() string {
	return "with §ActionStringProperty §StringComparator §StringValue"
}

type ActionTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (a ActionTagCondition) Validate() error {
	return a.ValidateStruct()
}

func (a ActionTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type EntityIntCondition struct {
	EntityIntProperty EntityIntProperty
	IntValue          SimpleIntValue
	IntComparator     IntComparator
}

func (e EntityIntCondition) Validate() error {
	return e.ValidateStruct()
}

func (e EntityIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityIntCondition) InteractionText() string {
	return "with §EntityIntProperty §IntComparator §IntValue"
}

type EntityStringCondition struct {
	EntityStringProperty EntityStringProperty
	StringValue          SimpleStringValue
	StringComparator     StringComparator
}

func (e EntityStringCondition) Validate() error {
	return e.ValidateStruct()
}

func (e EntityStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityStringCondition) InteractionText() string {
	return "with §EntityStringProperty §StringComparator §StringValue"
}

type EntityTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (e EntityTagCondition) Validate() error {
	return e.ValidateStruct()
}

func (e EntityTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type HeadquarterIntCondition struct {
	HeadquarterIntProperty HeadquarterIntProperty
	IntValue               SimpleIntValue
	IntComparator          IntComparator
}

func (h HeadquarterIntCondition) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterIntCondition) InteractionText() string {
	return "with §HeadquarterIntProperty §IntComparator §IntValue"
}

type HeadquarterStringCondition struct {
	HeadquarterStringProperty HeadquarterStringProperty
	StringValue               SimpleStringValue
	StringComparator          StringComparator
}

func (h HeadquarterStringCondition) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterStringCondition) InteractionText() string {
	return "with §HeadquarterStringProperty §StringComparator §StringValue"
}

type HeadquarterTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (h HeadquarterTagCondition) Validate() error {
	return h.ValidateStruct()
}

func (h HeadquarterTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h HeadquarterTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type PlaceIntCondition struct {
	PlaceIntProperty PlaceIntProperty
	IntValue         SimpleIntValue
	IntComparator    IntComparator
}

func (p PlaceIntCondition) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceIntCondition) InteractionText() string {
	return "with §PlaceIntProperty §IntComparator §IntValue"
}

type PlaceStringCondition struct {
	PlaceStringProperty PlaceStringProperty
	StringValue         SimpleStringValue
	StringComparator    StringComparator
}

func (p PlaceStringCondition) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceStringCondition) InteractionText() string {
	return "with §PlaceStringProperty §StringComparator §StringValue"
}

type PlaceTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (p PlaceTagCondition) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type PlayerIntCondition struct {
	PlayerIntProperty PlayerIntProperty
	IntValue          SimpleIntValue
	IntComparator     IntComparator
}

func (p PlayerIntCondition) Validate() error {
	return p.ValidateStruct()
}

func (p PlayerIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlayerIntCondition) InteractionText() string {
	return "with §PlayerIntProperty §IntComparator §IntValue"
}

package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type CardConditions struct {
	ActionConditions      *ActionConditions      `json:",omitempty"`
	EntityConditions      *EntityConditions      `json:",omitempty"`
	HeadquarterConditions *HeadquarterConditions `json:",omitempty"`
	PlaceConditions       *PlaceConditions       `json:",omitempty"`
	ThisConditions        *ThisCondition         `json:",omitempty"`
}

func (c CardConditions) ValidateType(r jsonschema.RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (c CardConditions) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(c)
}

type ActionConditions []ActionCondition

func (a ActionConditions) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range a {
		err := v.ValidateType(r)
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

func (e EntityConditions) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range e {
		err := v.ValidateType(r)
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

func (h HeadquarterConditions) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range h {
		err := v.ValidateType(r)
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

func (p PlaceConditions) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range p {
		err := v.ValidateType(r)
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

func (p PlayerConditions) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range p {
		err := v.ValidateType(r)
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

func (a ActionCondition) ValidateType(r jsonschema.RootElement) error {
	implementer, err := a.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (a ActionCondition) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(a)
}

type EntityCondition struct {
	EntityIntCondition    *EntityIntCondition    `json:",omitempty"`
	EntityStringCondition *EntityStringCondition `json:",omitempty"`
	EntityTagCondition    *EntityTagCondition    `json:",omitempty"`
}

func (e EntityCondition) ValidateType(r jsonschema.RootElement) error {
	implementer, err := e.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (e EntityCondition) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(e)
}

type HeadquarterCondition struct {
	HeadquarterIntCondition    *HeadquarterIntCondition    `json:",omitempty"`
	HeadquarterStringCondition *HeadquarterStringCondition `json:",omitempty"`
	HeadquarterTagCondition    *HeadquarterTagCondition    `json:",omitempty"`
}

func (h HeadquarterCondition) ValidateType(r jsonschema.RootElement) error {
	implementer, err := h.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (h HeadquarterCondition) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(h)
}

type PlaceCondition struct {
	PlaceIntCondition    *PlaceIntCondition    `json:",omitempty"`
	PlaceStringCondition *PlaceStringCondition `json:",omitempty"`
	PlaceTagCondition    *PlaceTagCondition    `json:",omitempty"`
}

func (p PlaceCondition) ValidateType(r jsonschema.RootElement) error {
	implementer, err := p.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (p PlaceCondition) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(p)
}

type PlayerCondition struct {
	PlayerIntCondition *PlayerIntCondition `json:",omitempty"`
}

func (p PlayerCondition) ValidateType(r jsonschema.RootElement) error {
	implementer, err := p.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (p PlayerCondition) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(p)
}

type ThisCondition struct{}

func (t ThisCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t ThisCondition) InteractionText() string {
	return "this"
}

type ActionIntCondition struct {
	ActionIntProperty ActionIntProperty
	IntValue          SimpleIntValue
	IntComparator     IntComparator
}

func (a ActionIntCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionIntCondition) InteractionText() string {
	return "with §ActionIntProperty §IntComparator §IntValue"
}

type ActionStringCondition struct {
	ActionStringProperty ActionStringProperty
	StringValue          SimpleStringValue
	StringComparator     StringComparator
}

func (a ActionStringCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionStringCondition) InteractionText() string {
	return "with §ActionStringProperty §StringComparator §StringValue"
}

type ActionTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (a ActionTagCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type EntityIntCondition struct {
	EntityIntProperty EntityIntProperty
	IntValue          SimpleIntValue
	IntComparator     IntComparator
}

func (e EntityIntCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityIntCondition) InteractionText() string {
	return "with §EntityIntProperty §IntComparator §IntValue"
}

type EntityStringCondition struct {
	EntityStringProperty EntityStringProperty
	StringValue          SimpleStringValue
	StringComparator     StringComparator
}

func (e EntityStringCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityStringCondition) InteractionText() string {
	return "with §EntityStringProperty §StringComparator §StringValue"
}

type EntityTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (e EntityTagCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type HeadquarterIntCondition struct {
	HeadquarterIntProperty HeadquarterIntProperty
	IntValue               SimpleIntValue
	IntComparator          IntComparator
}

func (h HeadquarterIntCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterIntCondition) InteractionText() string {
	return "with §HeadquarterIntProperty §IntComparator §IntValue"
}

type HeadquarterStringCondition struct {
	HeadquarterStringProperty HeadquarterStringProperty
	StringValue               SimpleStringValue
	StringComparator          StringComparator
}

func (h HeadquarterStringCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterStringCondition) InteractionText() string {
	return "with §HeadquarterStringProperty §StringComparator §StringValue"
}

type HeadquarterTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (h HeadquarterTagCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type PlaceIntCondition struct {
	PlaceIntProperty PlaceIntProperty
	IntValue         SimpleIntValue
	IntComparator    IntComparator
}

func (p PlaceIntCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceIntCondition) InteractionText() string {
	return "with §PlaceIntProperty §IntComparator §IntValue"
}

type PlaceStringCondition struct {
	PlaceStringProperty PlaceStringProperty
	StringValue         SimpleStringValue
	StringComparator    StringComparator
}

func (p PlaceStringCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceStringCondition) InteractionText() string {
	return "with §PlaceStringProperty §StringComparator §StringValue"
}

type PlaceTagCondition struct {
	StringValue      SimpleStringValue
	StringComparator StringComparator
}

func (p PlaceTagCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceTagCondition) InteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type PlayerIntCondition struct {
	PlayerIntProperty PlayerIntProperty
	IntValue          SimpleIntValue
	IntComparator     IntComparator
}

func (p PlayerIntCondition) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlayerIntCondition) InteractionText() string {
	return "with §PlayerIntProperty §IntComparator §IntValue"
}

package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type CardExtractors struct {
	ActionExtractors      *ActionExtractors      `json:",omitempty"`
	EntityExtractors      *EntityExtractors      `json:",omitempty"`
	HeadquarterExtractors *HeadquarterExtractors `json:",omitempty"`
	PlaceExtractors       *PlaceExtractors       `json:",omitempty"`
}

func (c CardExtractors) ValidateType(r jsonschema.RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (c CardExtractors) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(c)
}

type ActionExtractors []ActionExtractor

func (a ActionExtractors) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range a {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a ActionExtractors) MinMaxItems() (int, int) {
	return 0, 3
}

func (a ActionExtractors) ItemName() string {
	return jsonschema.ItemNameFromArray(a)
}

type EntityExtractors []EntityExtractor

func (e EntityExtractors) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range e {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e EntityExtractors) MinMaxItems() (int, int) {
	return 0, 3
}

func (e EntityExtractors) ItemName() string {
	return jsonschema.ItemNameFromArray(e)
}

type HeadquarterExtractors []HeadquarterExtractor

func (h HeadquarterExtractors) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range h {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (h HeadquarterExtractors) MinMaxItems() (int, int) {
	return 0, 3
}

func (h HeadquarterExtractors) ItemName() string {
	return jsonschema.ItemNameFromArray(h)
}

type PlaceExtractors []PlaceExtractor

func (p PlaceExtractors) ValidateType(r jsonschema.RootElement) error {
	errorRange := []error{}
	for _, v := range p {
		err := v.ValidateType(r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p PlaceExtractors) MinMaxItems() (int, int) {
	return 0, 3
}

func (p PlaceExtractors) ItemName() string {
	return jsonschema.ItemNameFromArray(p)
}

type ActionExtractor struct {
	ActionIntExtractor    *ActionIntExtractor    `json:",omitempty"`
	ActionStringExtractor *ActionStringExtractor `json:",omitempty"`
	ActionTargetExtractor *TargetExtractor       `json:",omitempty"`
}

func (a ActionExtractor) ValidateType(r jsonschema.RootElement) error {
	implementer, err := a.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (a ActionExtractor) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(a)
}

type EntityExtractor struct {
	EntityIntExtractor    *EntityIntExtractor    `json:",omitempty"`
	EntityStringExtractor *EntityStringExtractor `json:",omitempty"`
	EntityTargetExtractor *TargetExtractor       `json:",omitempty"`
}

func (e EntityExtractor) ValidateType(r jsonschema.RootElement) error {
	implementer, err := e.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (e EntityExtractor) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(e)
}

type HeadquarterExtractor struct {
	HeadquarterIntExtractor    *HeadquarterIntExtractor    `json:",omitempty"`
	HeadquarterStringExtractor *HeadquarterStringExtractor `json:",omitempty"`
	HeadquarterTargetExtractor *TargetExtractor            `json:",omitempty"`
}

func (h HeadquarterExtractor) ValidateType(r jsonschema.RootElement) error {
	implementer, err := h.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (h HeadquarterExtractor) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(h)
}

type PlaceExtractor struct {
	PlaceIntExtractor    *PlaceIntExtractor    `json:",omitempty"`
	PlaceStringExtractor *PlaceStringExtractor `json:",omitempty"`
	PlaceTargetExtractor *TargetExtractor      `json:",omitempty"`
}

func (p PlaceExtractor) ValidateType(r jsonschema.RootElement) error {
	implementer, err := p.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (p PlaceExtractor) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(p)
}

type ActionIntExtractor struct {
	ExtractIntProperty ActionIntProperty
	IntVariableName    IntVariableName
}

func (a ActionIntExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionIntExtractor) InteractionText() string {
	return "Set §IntVariableName to the actions §ExtractIntProperty."
}

type ActionStringExtractor struct {
	ExtractStringProperty ActionStringProperty
	StringVariableName    StringVariableName
}

func (a ActionStringExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionStringExtractor) InteractionText() string {
	return "Set §StringVariableName to the actions §ExtractStringProperty."
}

type EntityIntExtractor struct {
	ExtractIntProperty EntityIntProperty
	IntVariableName    IntVariableName
}

func (e EntityIntExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityIntExtractor) InteractionText() string {
	return "Set §IntVariableName to the entities §ExtractIntProperty."
}

type EntityStringExtractor struct {
	ExtractStringProperty EntityStringProperty
	StringVariableName    StringVariableName
}

func (e EntityStringExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityStringExtractor) InteractionText() string {
	return "Set §StringVariableName to the entities §ExtractStringProperty."
}

type HeadquarterIntExtractor struct {
	ExtractIntProperty HeadquarterIntProperty
	IntVariableName    IntVariableName
}

func (h HeadquarterIntExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterIntExtractor) InteractionText() string {
	return "Set §IntVariableName to the actions §ExtractIntProperty."
}

type HeadquarterStringExtractor struct {
	ExtractStringProperty HeadquarterStringProperty
	StringVariableName    StringVariableName
}

func (h HeadquarterStringExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h HeadquarterStringExtractor) InteractionText() string {
	return "Set §StringVariableName to the actions §ExtractStringProperty."
}

type PlaceIntExtractor struct {
	ExtractIntProperty PlaceIntProperty
	IntVariableName    IntVariableName
}

func (p PlaceIntExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceIntExtractor) InteractionText() string {
	return "Set §IntVariableName to the places §ExtractIntProperty."
}

type PlaceStringExtractor struct {
	ExtractStringProperty PlaceStringProperty
	StringVariableName    StringVariableName
}

func (p PlaceStringExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceStringExtractor) InteractionText() string {
	return "Set §StringVariableName to the places §ExtractStringProperty."
}

type TargetExtractor struct {
	TargetVariableName TargetVariableName
}

func (t TargetExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t TargetExtractor) InteractionText() string {
	return "That card is marked as §TargetVariableName."
}

type IntExtractor struct {
	IntVariableName IntVariableName
}

func (i IntExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(i, r)
}

func (i IntExtractor) InteractionText() string {
	return "Set §IntVariableName."
}

type StringExtractor struct {
	StringVariableName StringVariableName
}

func (s StringExtractor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s StringExtractor) InteractionText() string {
	return "Set §StringVariableName."
}

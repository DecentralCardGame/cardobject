package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type CardExtractors struct {
	ActionExtractors *ActionExtractors `json:",omitempty"`
	EntityExtractors *EntityExtractors `json:",omitempty"`
	PlaceExtractors  *PlaceExtractors  `json:",omitempty"`
}

func (c CardExtractors) Validate() error {
	return c.ValidateInterface()
}

func (c CardExtractors) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type ActionExtractors []ActionExtractor

func (a ActionExtractors) Validate() error {
	return a.ValidateArray()
}

func (a ActionExtractors) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a ActionExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (a ActionExtractors) GetItemName() string {
	return jsonschema.GetItemNameFromArray(a)
}

type EntityExtractors []EntityExtractor

func (e EntityExtractors) Validate() error {
	return e.ValidateArray()
}

func (e EntityExtractors) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e EntityExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (e EntityExtractors) GetItemName() string {
	return jsonschema.GetItemNameFromArray(e)
}

type PlaceExtractors []PlaceExtractor

func (p PlaceExtractors) Validate() error {
	return p.ValidateArray()
}

func (p PlaceExtractors) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p PlaceExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (p PlaceExtractors) GetItemName() string {
	return jsonschema.GetItemNameFromArray(p)
}

type ActionExtractor struct {
	ActionIntExtractor    *ActionIntExtractor    `json:",omitempty"`
	ActionStringExtractor *ActionStringExtractor `json:",omitempty"`
	ActionTargetExtractor *TargetExtractor       `json:",omitempty"`
}

func (a ActionExtractor) Validate() error {
	return a.ValidateInterface()
}

func (a ActionExtractor) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type EntityExtractor struct {
	EntityIntExtractor    *EntityIntExtractor    `json:",omitempty"`
	EntityStringExtractor *EntityStringExtractor `json:",omitempty"`
	EntityTargetExtractor *TargetExtractor       `json:",omitempty"`
}

func (e EntityExtractor) Validate() error {
	return e.ValidateInterface()
}

func (e EntityExtractor) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type PlaceExtractor struct {
	PlaceIntExtractor    *PlaceIntExtractor    `json:",omitempty"`
	PlaceStringExtractor *PlaceStringExtractor `json:",omitempty"`
	PlaceTargetExtractor *TargetExtractor      `json:",omitempty"`
}

func (p PlaceExtractor) Validate() error {
	return p.ValidateInterface()
}

func (p PlaceExtractor) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type ActionIntExtractor struct {
	ExtractIntProperty ActionIntProperty
	IntVariableName    IntVariableName
}

func (a ActionIntExtractor) Validate() error {
	return a.ValidateStruct()
}

func (a ActionIntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionIntExtractor) GetInteractionText() string {
	return "Set §IntVariableName to the actions §ExtractIntProperty."
}

type ActionStringExtractor struct {
	ExtractStringProperty ActionStringProperty
	StringVariableName    StringVariableName
}

func (a ActionStringExtractor) Validate() error {
	return a.ValidateStruct()
}

func (a ActionStringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionStringExtractor) GetInteractionText() string {
	return "Set §StringVariableName to the actionss §ExtractStringProperty."
}

type EntityIntExtractor struct {
	ExtractIntProperty EntityIntProperty
	IntVariableName    IntVariableName
}

func (e EntityIntExtractor) Validate() error {
	return e.ValidateStruct()
}

func (e EntityIntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityIntExtractor) GetInteractionText() string {
	return "Set §IntVariableName to the entities §ExtractIntProperty."
}

type EntityStringExtractor struct {
	ExtractStringProperty EntityStringProperty
	StringVariableName    StringVariableName
}

func (e EntityStringExtractor) Validate() error {
	return e.ValidateStruct()
}

func (e EntityStringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityStringExtractor) GetInteractionText() string {
	return "Set §StringVariableName to the entities §ExtractStringProperty."
}

type PlaceIntExtractor struct {
	ExtractIntProperty PlaceIntProperty
	IntVariableName    IntVariableName
}

func (p PlaceIntExtractor) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceIntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceIntExtractor) GetInteractionText() string {
	return "Set §IntVariableName to the places §ExtractIntProperty."
}

type PlaceStringExtractor struct {
	ExtractStringProperty PlaceStringProperty
	StringVariableName    StringVariableName
}

func (p PlaceStringExtractor) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceStringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceStringExtractor) GetInteractionText() string {
	return "Set §StringVariableName to the places §ExtractStringProperty."
}

type TargetExtractor struct {
	TargetVariableName TargetVariableName
}

func (t TargetExtractor) Validate() error {
	return t.ValidateStruct()
}

func (t TargetExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t TargetExtractor) GetInteractionText() string {
	return "That card is marked as §TargetVariableName."
}

type IntExtractor struct {
	IntVariableName IntVariableName
}

func (i IntExtractor) Validate() error {
	return i.ValidateStruct()
}

func (i IntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(i)
}

func (i IntExtractor) GetInteractionText() string {
	return "Set §IntVariableName."
}

type StringExtractor struct {
	StringVariableName StringVariableName
}

func (s StringExtractor) Validate() error {
	return s.ValidateStruct()
}

func (s StringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s StringExtractor) GetInteractionText() string {
	return "Set §StringVariableName."
}

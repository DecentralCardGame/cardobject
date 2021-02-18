package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type cardExtractors struct {
	ActionExtractors *actionExtractors `json:",omitempty"`
	EntityExtractors *entityExtractors `json:",omitempty"`
	PlaceExtractors  *placeExtractors  `json:",omitempty"`
}

func (c cardExtractors) Validate() error {
	return c.ValidateInterface()
}

func (c cardExtractors) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type actionExtractors []actionExtractor

func (a actionExtractors) Validate() error {
	return a.ValidateArray()
}

func (a actionExtractors) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a actionExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (a actionExtractors) GetItemName() string {
	return jsonschema.GetItemNameFromArray(a)
}

type entityExtractors []entityExtractor

func (e entityExtractors) Validate() error {
	return e.ValidateArray()
}

func (e entityExtractors) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e entityExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (e entityExtractors) GetItemName() string {
	return jsonschema.GetItemNameFromArray(e)
}

type placeExtractors []placeExtractor

func (p placeExtractors) Validate() error {
	return p.ValidateArray()
}

func (p placeExtractors) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p placeExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (p placeExtractors) GetItemName() string {
	return jsonschema.GetItemNameFromArray(p)
}

type actionExtractor struct {
	ActionIntExtractor    *actionIntExtractor    `json:",omitempty"`
	ActionStringExtractor *actionStringExtractor `json:",omitempty"`
	ActionTargetExtractor *targetExtractor       `json:",omitempty"`
}

func (a actionExtractor) Validate() error {
	return a.ValidateInterface()
}

func (a actionExtractor) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type entityExtractor struct {
	EntityIntExtractor    *entityIntExtractor    `json:",omitempty"`
	EntityStringExtractor *entityStringExtractor `json:",omitempty"`
	EntityTargetExtractor *targetExtractor       `json:",omitempty"`
}

func (e entityExtractor) Validate() error {
	return e.ValidateInterface()
}

func (e entityExtractor) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type placeExtractor struct {
	PlaceIntExtractor    *placeIntExtractor    `json:",omitempty"`
	PlaceStringExtractor *placeStringExtractor `json:",omitempty"`
	PlaceTargetExtractor *targetExtractor      `json:",omitempty"`
}

func (p placeExtractor) Validate() error {
	return p.ValidateInterface()
}

func (p placeExtractor) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type actionIntExtractor struct {
	ExtractIntProperty actionIntProperty
	IntVariableName    intVariableName
}

func (a actionIntExtractor) Validate() error {
	return a.ValidateStruct()
}

func (a actionIntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionIntExtractor) GetInteractionText() string {
	return "Set §IntVariableName to the actions §ExtractIntProperty."
}

type actionStringExtractor struct {
	ExtractStringProperty actionStringProperty
	StringVariableName    stringVariableName
}

func (a actionStringExtractor) Validate() error {
	return a.ValidateStruct()
}

func (a actionStringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionStringExtractor) GetInteractionText() string {
	return "Set §StringVariableName to the actionss §ExtractStringProperty."
}

type entityIntExtractor struct {
	ExtractIntProperty entityIntProperty
	IntVariableName    intVariableName
}

func (e entityIntExtractor) Validate() error {
	return e.ValidateStruct()
}

func (e entityIntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityIntExtractor) GetInteractionText() string {
	return "Set §IntVariableName to the entities §ExtractIntProperty."
}

type entityStringExtractor struct {
	ExtractStringProperty entityStringProperty
	StringVariableName    stringVariableName
}

func (e entityStringExtractor) Validate() error {
	return e.ValidateStruct()
}

func (e entityStringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityStringExtractor) GetInteractionText() string {
	return "Set §StringVariableName to the entities §ExtractStringProperty."
}

type placeIntExtractor struct {
	ExtractIntProperty placeIntProperty
	IntVariableName    intVariableName
}

func (p placeIntExtractor) Validate() error {
	return p.ValidateStruct()
}

func (p placeIntExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeIntExtractor) GetInteractionText() string {
	return "Set §IntVariableName to the places §ExtractIntProperty."
}

type placeStringExtractor struct {
	ExtractStringProperty placeStringProperty
	StringVariableName    stringVariableName
}

func (p placeStringExtractor) Validate() error {
	return p.ValidateStruct()
}

func (p placeStringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeStringExtractor) GetInteractionText() string {
	return "Set §StringVariableName to the places §ExtractStringProperty."
}

type targetExtractor struct {
	TargetVariableName targetVariableName
}

func (t targetExtractor) Validate() error {
	return t.ValidateStruct()
}

func (t targetExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t targetExtractor) GetInteractionText() string {
	return "That card is marked as §TargetVariableName."
}

type intExtractor struct {
	IntVariableName intVariableName
}

func (i intExtractor) Validate() error {
	return i.ValidateStruct()
}

func (i intExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(i)
}

func (i intExtractor) GetInteractionText() string {
	return "Set §IntVariableName."
}

type stringExtractor struct {
	StringVariableName stringVariableName
}

func (s stringExtractor) Validate() error {
	return s.ValidateStruct()
}

func (s stringExtractor) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s stringExtractor) GetInteractionText() string {
	return "Set §StringVariableName."
}

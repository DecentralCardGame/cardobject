package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type cardExtractors struct {
	*jsonschema.BasicInterface
	ActionExtractors *actionExtractors `json:",omitempty"`
	EntityExtractors *entityExtractors `json:",omitempty"`
	PlaceExtractors  *placeExtractors  `json:",omitempty"`
}

type actionExtractors []actionExtractor

func (a actionExtractors) Validate() error {
	return a.ValidateArray()
}

func (a actionExtractors) ValidateArray() error {
	return nil
}

func (a actionExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

type entityExtractors []entityExtractor

func (e entityExtractors) Validate() error {
	return e.ValidateArray()
}

func (e entityExtractors) ValidateArray() error {
	return nil
}

func (e entityExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

type placeExtractors []placeExtractor

func (p placeExtractors) Validate() error {
	return p.ValidateArray()
}

func (p placeExtractors) ValidateArray() error {
	return nil
}

func (p placeExtractors) GetMinMaxItems() (int, int) {
	return 0, 3
}

type actionExtractor struct {
	*jsonschema.BasicInterface
	ActionIntExtractor    *actionIntExtractor    `json:",omitempty"`
	ActionStringExtractor *actionStringExtractor `json:",omitempty"`
	ActionTargetExtractor *targetExtractor       `json:",omitempty"`
}

type entityExtractor struct {
	*jsonschema.BasicInterface
	EntityIntExtractor    *entityIntExtractor    `json:",omitempty"`
	EntityStringExtractor *entityStringExtractor `json:",omitempty"`
	EntityTargetExtractor *targetExtractor       `json:",omitempty"`
}

type placeExtractor struct {
	*jsonschema.BasicInterface
	PlaceIntExtractor    *placeIntExtractor    `json:",omitempty"`
	PlaceStringExtractor *placeStringExtractor `json:",omitempty"`
	PlaceTargetExtractor *targetExtractor      `json:",omitempty"`
}

type actionIntExtractor struct {
	*jsonschema.BasicStruct
	ExtractIntProperty actionIntProperty
	IntVariableName    intVariableName
}

func (a actionIntExtractor) GetInteractionText() string {
	return ""
}

type actionStringExtractor struct {
	*jsonschema.BasicStruct
	ExtractStringProperty actionStringProperty
	StringVariableName    stringVariableName
}

func (a actionStringExtractor) GetInteractionText() string {
	return ""
}

type entityIntExtractor struct {
	*jsonschema.BasicStruct
	ExtractIntProperty entityIntProperty
	IntVariableName    intVariableName
}

func (e entityIntExtractor) GetInteractionText() string {
	return ""
}

type entityStringExtractor struct {
	*jsonschema.BasicStruct
	ExtractStringProperty entityStringProperty
	StringVariableName    stringVariableName
}

func (e entityStringExtractor) GetInteractionText() string {
	return ""
}

type placeIntExtractor struct {
	*jsonschema.BasicStruct
	ExtractIntProperty placeIntProperty
	IntVariableName    intVariableName
}

func (p placeIntExtractor) GetInteractionText() string {
	return ""
}

type placeStringExtractor struct {
	*jsonschema.BasicStruct
	ExtractStringProperty placeStringProperty
	StringVariableName    stringVariableName
}

func (p placeStringExtractor) GetInteractionText() string {
	return ""
}

type targetExtractor struct {
	*jsonschema.BasicStruct
	TargetVariableName targetVariableName
}

func (t targetExtractor) GetInteractionText() string {
	return ""
}

type intExtractor struct {
	*jsonschema.BasicStruct
	IntVariableName intVariableName
}

func (i intExtractor) GetInteractionText() string {
	return ""
}

type stringExtractor struct {
	*jsonschema.BasicStruct
	StringVariableName stringVariableName
}

func (s stringExtractor) GetInteractionText() string {
	return ""
}

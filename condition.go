package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type cardConditions struct {
	*jsonschema.BasicInterface
	ActionConditions *actionConditions `json:",omitempty"`
	EntityConditions *entityConditions `json:",omitempty"`
	PlaceConditions  *placeConditions  `json:",omitempty"`
	ThisConditions   *thisCondition    `json:",omitempty"`
}

type actionConditions []actionCondition

func (a actionConditions) Validate() error {
	return a.ValidateArray()
}

func (a actionConditions) ValidateArray() error {
	return nil
}

func (a actionConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

type entityConditions []entityCondition

func (e entityConditions) Validate() error {
	return e.ValidateArray()
}

func (e entityConditions) ValidateArray() error {
	return nil
}

func (e entityConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

type placeConditions []placeCondition

func (p placeConditions) Validate() error {
	return p.ValidateArray()
}

func (p placeConditions) ValidateArray() error {
	return nil
}

func (p placeConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

type playerConditions []playerCondition

func (p playerConditions) Validate() error {
	return p.ValidateArray()
}

func (p playerConditions) ValidateArray() error {
	return nil
}

func (p playerConditions) GetMinMayItems() (int, int) {
	return 0, 3
}

type actionCondition struct {
	*jsonschema.BasicInterface
	ActionIntCondition    *actionIntCondition    `json:",omitempty"`
	ActionStringCondition *actionStringCondition `json:",omitempty"`
	ActionTagCondition    *actionTagCondition    `json:",omitempty"`
}

type entityCondition struct {
	*jsonschema.BasicInterface
	EntityIntCondition    *entityIntCondition    `json:",omitempty"`
	EntityStringCondition *entityStringCondition `json:",omitempty"`
	EntityTagCondition    *entityTagCondition    `json:",omitempty"`
}

type placeCondition struct {
	*jsonschema.BasicInterface
	PlaceIntCondition    *placeIntCondition    `json:",omitempty"`
	PlaceStringCondition *placeStringCondition `json:",omitempty"`
	PlaceTagCondition    *placeTagCondition    `json:",omitempty"`
}

type playerCondition struct {
	*jsonschema.BasicInterface
	PlayerIntCondition *playerIntCondition `json:",omitempty"`
}

type thisCondition struct {
	*jsonschema.BasicStruct
}

func (t thisCondition) GetInteractionText() string {
	return ""
}

type actionIntCondition struct {
	*jsonschema.BasicStruct
	ActionIntProperty actionIntProperty
	IntValue          simpleIntValue
	IntComparator     intComparator
}

func (a actionIntCondition) GetInteractionText() string {
	return ""
}

type actionStringCondition struct {
	*jsonschema.BasicStruct
	ActionStringProperty actionStringProperty
	StringValue          simpleStringValue
	StringComparator     stringComparator
}

func (a actionStringCondition) GetInteractionText() string {
	return ""
}

type actionTagCondition struct {
	*jsonschema.BasicStruct
	StringValue      simpleStringValue
	StringComparator stringComparator
}

func (a actionTagCondition) GetInteractionText() string {
	return ""
}

type entityIntCondition struct {
	*jsonschema.BasicStruct
	EntityIntProperty entityIntProperty
	IntValue          simpleIntValue
	IntComparator     intComparator
}

func (e entityIntCondition) GetInteractionText() string {
	return ""
}

type entityStringCondition struct {
	*jsonschema.BasicStruct
	EntityStringProperty entityStringProperty
	StringValue          simpleStringValue
	StringComparator     stringComparator
}

func (e entityStringCondition) GetInteractionText() string {
	return ""
}

type entityTagCondition struct {
	*jsonschema.BasicStruct
	StringValue      simpleStringValue
	StringComparator stringComparator
}

func (e entityTagCondition) GetInteractionText() string {
	return ""
}

type placeIntCondition struct {
	*jsonschema.BasicStruct
	PlaceIntProperty placeIntProperty
	IntValue         simpleIntValue
	IntComparator    intComparator
}

func (p placeIntCondition) GetInteractionText() string {
	return ""
}

type placeStringCondition struct {
	*jsonschema.BasicStruct
	PlaceStringProperty placeStringProperty
	StringValue         simpleStringValue
	StringComparator    stringComparator
}

func (p placeStringCondition) GetInteractionText() string {
	return ""
}

type placeTagCondition struct {
	*jsonschema.BasicStruct
	StringValue      simpleStringValue
	StringComparator stringComparator
}

func (p placeTagCondition) GetInteractionText() string {
	return ""
}

type playerIntCondition struct {
	*jsonschema.BasicStruct
	PlayerIntProperty playerIntProperty
	IntValue          simpleIntValue
	IntComparator     intComparator
}

func (p playerIntCondition) GetInteractionText() string {
	return ""
}

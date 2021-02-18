package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type cardConditions struct {
	ActionConditions *actionConditions `json:",omitempty"`
	EntityConditions *entityConditions `json:",omitempty"`
	PlaceConditions  *placeConditions  `json:",omitempty"`
	ThisConditions   *thisCondition    `json:",omitempty"`
}

func (c cardConditions) Validate() error {
	return c.ValidateInterface()
}

func (c cardConditions) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type actionConditions []actionCondition

func (a actionConditions) Validate() error {
	return a.ValidateArray()
}

func (a actionConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range a {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (a actionConditions) GetItemName() string {
	return jsonschema.GetItemNameFromArray(a)
}

func (a actionConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

type entityConditions []entityCondition

func (e entityConditions) Validate() error {
	return e.ValidateArray()
}

func (e entityConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range e {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (e entityConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (e entityConditions) GetItemName() string {
	return jsonschema.GetItemNameFromArray(e)
}

type placeConditions []placeCondition

func (p placeConditions) Validate() error {
	return p.ValidateArray()
}

func (p placeConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p placeConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (p placeConditions) GetItemName() string {
	return jsonschema.GetItemNameFromArray(p)
}

type playerConditions []playerCondition

func (p playerConditions) Validate() error {
	return p.ValidateArray()
}

func (p playerConditions) ValidateArray() error {
	errorRange := []error{}
	for _, v := range p {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return jsonschema.CombineErrors(errorRange)
}

func (p playerConditions) GetMinMaxItems() (int, int) {
	return 0, 3
}

func (p playerConditions) GetItemName() string {
	return jsonschema.GetItemNameFromArray(p)
}

type actionCondition struct {
	ActionIntCondition    *actionIntCondition    `json:",omitempty"`
	ActionStringCondition *actionStringCondition `json:",omitempty"`
	ActionTagCondition    *actionTagCondition    `json:",omitempty"`
}

func (a actionCondition) Validate() error {
	return a.ValidateInterface()
}

func (a actionCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(a)
}

type entityCondition struct {
	EntityIntCondition    *entityIntCondition    `json:",omitempty"`
	EntityStringCondition *entityStringCondition `json:",omitempty"`
	EntityTagCondition    *entityTagCondition    `json:",omitempty"`
}

func (e entityCondition) Validate() error {
	return e.ValidateInterface()
}

func (e entityCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type placeCondition struct {
	PlaceIntCondition    *placeIntCondition    `json:",omitempty"`
	PlaceStringCondition *placeStringCondition `json:",omitempty"`
	PlaceTagCondition    *placeTagCondition    `json:",omitempty"`
}

func (p placeCondition) Validate() error {
	return p.ValidateInterface()
}

func (p placeCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type playerCondition struct {
	PlayerIntCondition *playerIntCondition `json:",omitempty"`
}

func (p playerCondition) Validate() error {
	return p.ValidateInterface()
}

func (p playerCondition) ValidateInterface() error {
	return jsonschema.ValidateInterface(p)
}

type thisCondition struct{}

func (t thisCondition) Validate() error {
	return t.ValidateStruct()
}

func (t thisCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t thisCondition) GetInteractionText() string {
	return "this"
}

type actionIntCondition struct {
	ActionIntProperty actionIntProperty
	IntValue          simpleIntValue
	IntComparator     intComparator
}

func (a actionIntCondition) Validate() error {
	return a.ValidateStruct()
}

func (a actionIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionIntCondition) GetInteractionText() string {
	return "with §ActionIntProperty §IntComparator §IntValue"
}

type actionStringCondition struct {
	ActionStringProperty actionStringProperty
	StringValue          simpleStringValue
	StringComparator     stringComparator
}

func (a actionStringCondition) Validate() error {
	return a.ValidateStruct()
}

func (a actionStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionStringCondition) GetInteractionText() string {
	return "with §ActionStringProperty §StringComparator §StringValue"
}

type actionTagCondition struct {
	StringValue      simpleStringValue
	StringComparator stringComparator
}

func (a actionTagCondition) Validate() error {
	return a.ValidateStruct()
}

func (a actionTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionTagCondition) GetInteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type entityIntCondition struct {
	EntityIntProperty entityIntProperty
	IntValue          simpleIntValue
	IntComparator     intComparator
}

func (e entityIntCondition) Validate() error {
	return e.ValidateStruct()
}

func (e entityIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityIntCondition) GetInteractionText() string {
	return "with §EntityIntProperty §IntComparator §IntValue"
}

type entityStringCondition struct {
	EntityStringProperty entityStringProperty
	StringValue          simpleStringValue
	StringComparator     stringComparator
}

func (e entityStringCondition) Validate() error {
	return e.ValidateStruct()
}

func (e entityStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityStringCondition) GetInteractionText() string {
	return "with §EntityStringProperty §StringComparator §StringValue"
}

type entityTagCondition struct {
	StringValue      simpleStringValue
	StringComparator stringComparator
}

func (e entityTagCondition) Validate() error {
	return e.ValidateStruct()
}

func (e entityTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityTagCondition) GetInteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type placeIntCondition struct {
	PlaceIntProperty placeIntProperty
	IntValue         simpleIntValue
	IntComparator    intComparator
}

func (p placeIntCondition) Validate() error {
	return p.ValidateStruct()
}

func (p placeIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeIntCondition) GetInteractionText() string {
	return "with §PlaceIntProperty §IntComparator §IntValue"
}

type placeStringCondition struct {
	PlaceStringProperty placeStringProperty
	StringValue         simpleStringValue
	StringComparator    stringComparator
}

func (p placeStringCondition) Validate() error {
	return p.ValidateStruct()
}

func (p placeStringCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeStringCondition) GetInteractionText() string {
	return "with §PlaceStringProperty §StringComparator §StringValue"
}

type placeTagCondition struct {
	StringValue      simpleStringValue
	StringComparator stringComparator
}

func (p placeTagCondition) Validate() error {
	return p.ValidateStruct()
}

func (p placeTagCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeTagCondition) GetInteractionText() string {
	return "with tag §StringComparator §StringValue"
}

type playerIntCondition struct {
	PlayerIntProperty playerIntProperty
	IntValue          simpleIntValue
	IntComparator     intComparator
}

func (p playerIntCondition) Validate() error {
	return p.ValidateStruct()
}

func (p playerIntCondition) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p playerIntCondition) GetInteractionText() string {
	return "with §PlayerIntProperty §IntComparator §IntValue"
}

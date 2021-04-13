package jsonschema

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type cardObject struct {
	Action *action
}

func (c cardObject) Validate() error {
	return c.ValidateInterface()
}

func (c cardObject) ValidateInterface() error {
	return ValidateInterface(c)
}

type action struct {
	Name        name
	Costs       costs
	Time        *time
	MultipleUse multipleUse
}

func (a action) Classes() []string {
	return []string{"Technology", "Nature"}
}

func (a action) Validate() error {
	return a.ValidateStruct()
}

func (a action) ValidateStruct() error {
	return ValidateStruct(a)
}

func (a action) InteractionText() string {
	return "§Name §Costs §Time §MultipleUse"
}

type costs []cost

func (c costs) Validate() error {
	return c.ValidateArray()
}

func (c costs) ValidateArray() error {
	errorRange := []error{}
	for _, v := range c {
		err := v.Validate()
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return CombineErrors(errorRange)
}

func (c costs) MinMaxItems() (int, int) {
	return 1, 3
}

func (c costs) ItemName() string {
	return ItemNameFromArray(c)
}

type cost BasicEnum

func (c cost) Validate() error {
	return c.ValidateEnum()
}

func (c cost) ValidateEnum() error {
	return nil
}

func (c cost) EnumValues() []string {
	return []string{"eins", "zwei"}
}

type name BasicString

func (n name) Validate() error {
	return n.ValidateString()
}

func (n name) ValidateString() error {
	return nil
}

func (n name) MinMaxLength() (int, int) {
	return 2, 30
}

type time BasicInt

func (t time) Validate() error {
	return t.ValidateInt()
}

func (t time) ValidateInt() error {
	return nil
}

func (t time) MinMax() (int, int) {
	return 1, 32
}

type multipleUse BasicBool

func (m multipleUse) Validate() error {
	return m.ValidateBool()
}

func (m multipleUse) ValidateBool() error {
	return nil
}

func TestReflect(t *testing.T) {
	schema, reflectErr := Reflect(&cardObject{})
	if reflectErr != nil {
		t.Error(reflectErr)
	} else {
		json, err := schema.MarshalJSONForJS("Card")
		if err != nil {
			t.Error(err)
		}
		println(string(json))
	}
}

func TestValidate(t *testing.T) {
	var c cardObject
	dat, readErr := ioutil.ReadFile("testJsons/testCard.json")
	if readErr != nil {
		t.Error(readErr)
	}

	marshalErr := json.Unmarshal(dat, &c)
	if marshalErr != nil {
		t.Error(marshalErr)
	}

	//println(c.Action.Name)
	//validateErr := c.Action.Time.Validate
	validateErr := c.Validate()
	if validateErr != nil {
		t.Error(validateErr)
	}
}

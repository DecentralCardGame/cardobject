package jsonschema

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type cardObject struct {
	Action *action
}

func (c cardObject) CheckRootRequirements(s []string) error {
	return nil
}

func (c cardObject) Validate() error {
	return c.ValidateType(c)
}

func (c cardObject) ValidateType(r RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (c cardObject) FindImplementer() (Validateable, error) {
	return FindImplementer(c)
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

func (a action) ValidateType(r RootElement) error {
	return ValidateStruct(a, r)
}

func (a action) InteractionText() string {
	return "§Name §Costs §Time §MultipleUse"
}

func (a action) Description() string {
	return "This allows you to build an Action-card"
}

type costs []cost

func (c costs) ValidateType(r RootElement) error {
	errorRange := []error{}
	for _, v := range c {
		err := v.ValidateType(r)
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

func (c cost) ValidateType(r RootElement) error {
	return nil
}

func (c cost) EnumValues() []string {
	return []string{"eins", "zwei"}
}

type name BasicString

func (n name) ValidateType(r RootElement) error {
	return nil
}

func (n name) MinMaxLength() (int, int) {
	return 2, 30
}

type time BasicInt

func (t time) ValidateType(r RootElement) error {
	return nil
}

func (t time) MinMax() (int, int) {
	return 1, 32
}

type multipleUse BasicBool

func (m multipleUse) ValidateType(r RootElement) error {
	return nil
}

func (m multipleUse) Default() bool {
	return false
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

func TestValidateType(t *testing.T) {
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

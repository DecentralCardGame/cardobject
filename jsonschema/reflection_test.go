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

func (c cardObject) ValidateRoot() error {
	return c.Validate(c)
}

func (c cardObject) Validate(r RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.Validate(r)
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

func (a action) Validate(r RootElement) error {
	return ValidateStruct(a, r)
}

func (a action) InteractionText() string {
	return "§Name §Costs §Time §MultipleUse"
}

func (a action) Description() string {
	return "This allows you to build an Action-card"
}

type costs []cost

func (c costs) Validate(r RootElement) error {
	errorRange := []error{}
	for _, v := range c {
		err := v.Validate(r)
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

func (c cost) Validate(r RootElement) error {
	return nil
}

func (c cost) EnumValues() []string {
	return []string{"eins", "zwei"}
}

type name BasicString

func (n name) Validate(r RootElement) error {
	return nil
}

func (n name) MinMaxLength() (int, int) {
	return 2, 30
}

type time BasicInt

func (t time) Validate(r RootElement) error {
	return nil
}

func (t time) MinMax() (int, int) {
	return 1, 32
}

type multipleUse BasicBool

func (m multipleUse) Validate(r RootElement) error {
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
	validateErr := c.ValidateRoot()
	if validateErr != nil {
		t.Error(validateErr)
	}
}

package cardobject

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/DecentralCardGame/jsonschema"
)

func writeToFile(schema string) {
	f, err := os.Create("schema.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(schema)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCardSerializationAction1(t *testing.T) {
	schema, textErr := jsonschema.Reflect(&card{})
	if textErr != nil {
		t.Error(textErr)
	}
	json, err := schema.MarshalJSONForJS("Card")
	if err != nil {
		t.Error(err)
	}
	writeToFile(string(json))
}

func TestValidateAction(t *testing.T) {
	var c card
	dat, readErr := ioutil.ReadFile("testJsons/actionTest1.json")
	if readErr != nil {
		t.Error(readErr)
	}

	marshalErr := json.Unmarshal(dat, &c)
	if marshalErr != nil {
		t.Error(marshalErr)
	}
	validateErr := c.Validate()
	if validateErr != nil {
		t.Error(validateErr)
	}
}

func TestValidateEntity(t *testing.T) {
	var c card
	dat, readErr := ioutil.ReadFile("testJsons/entityTest1.json")
	if readErr != nil {
		t.Error(readErr)
	}

	marshalErr := json.Unmarshal(dat, &c)
	if marshalErr != nil {
		t.Error(marshalErr)
	}
	validateErr := c.Validate()
	if validateErr != nil {
		t.Error(validateErr)
	}
}

func TestValidatePlace(t *testing.T) {
	var c card
	dat, readErr := ioutil.ReadFile("testJsons/place1Test.json")
	if readErr != nil {
		t.Error(readErr)
	}

	marshalErr := json.Unmarshal(dat, &c)
	if marshalErr != nil {
		t.Error(marshalErr)
	}
	validateErr := c.Validate()
	if validateErr != nil {
		t.Error(validateErr)
	}
}

package cardobject

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"cardobject/jsonschema"
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
	schema, textErr := jsonschema.Reflect(&Card{})
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
	dat, readErr := ioutil.ReadFile("testJsons/actionTest1.json")
	if readErr != nil {
		t.Error(readErr)
	}

	_, err := NewCardFromJson(string(dat))

	if err != nil {
		t.Error(err)
	}
}

func TestValidateEntity(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/entityTest1.json")
	if readErr != nil {
		t.Error(readErr)
	}

	_, err := NewCardFromJson(string(dat))

	if err != nil {
		t.Error(err)
	}
}

func TestValidatePlace(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/place1Test.json")
	if readErr != nil {
		t.Error(readErr)
	}

	_, err := NewCardFromJson(string(dat))

	if err != nil {
		t.Error(err)
	}
}

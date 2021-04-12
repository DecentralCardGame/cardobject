package keywords

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

func writeToFile(schema string) {
	f, err := os.Create("schema.json")
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

func TestValidatePayArm(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/payArm.json")
	if readErr != nil {
		t.Error(readErr)
	}

	card, err := Unmarshal(dat)

	if err != nil {
		t.Error(err)
	}

	json, _ := json.Marshal(card)
	println(string(json))
}

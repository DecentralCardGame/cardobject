package cardobject

import (
	"fmt"
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
	json, err := schema.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	writeToFile(string(json))
}

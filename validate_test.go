package cardobject

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestCardSerializationAction1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entityTest1.json")
	input := string(file)
	card, err := NewCardFromJson(input)
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(card)
	if err != nil {
		t.Error(err)
	}
	println(string(b))
}

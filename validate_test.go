package cardobject

import (
	"io/ioutil"
	"testing"
)

func TestCardSerializationAction1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/place1Test.json")
	input := string(file)
	_, err := NewCardFromJson(input)
	if err != nil {
		t.Error(err)
	}
}

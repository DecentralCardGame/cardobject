package cardCreatorObject

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func FunctionalCardJson(cardJson string) (string, error) {
	var card card
	err := json.Unmarshal([]byte(cardJson), &card)
	if err != nil {
		return "Can't deserialize", err
	}
	bytes, err := json.Marshal(card)
	if err != nil {
		return "Can't serialize", err
	}
	return string(bytes), nil
}

func TestCardSerializationAction1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action1Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if input != output {
		t.Errorf("In- and output are not equal: " + output)
	}
}

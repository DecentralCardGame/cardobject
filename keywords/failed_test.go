package keywords

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestFailed(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/failedJson.json")
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

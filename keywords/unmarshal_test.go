package keywords

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"
)

func emptyTestCard() Card {
	return Card{}
}

func TestUnmarshaling(t *testing.T) {
	dat, _ := ioutil.ReadFile("testJsons/keywordedCard.json")

	card, err := Unmarshal(dat)
	if err != nil {
		t.Error(err)
	}

	cardJSON, _ := json.Marshal(card)

	input := string(dat)

	output := string(cardJSON)

	if input != output {
		t.Error(errors.New("ResolvedCard doesn't match the ground truth: " + output + " is not " + input))
	}
}

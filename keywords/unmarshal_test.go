package keywords

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"
)

func TestUnmarshaling(t *testing.T) {
	dat, _ := ioutil.ReadFile("testJsons/keywordedCard.json")
	groundTruth, _ := ioutil.ReadFile("testJsons/resolvedGroundTruth.json")

	card, err := Unmarshal(dat)
	if err != nil {
		t.Error(err)
	}

	resolvedCard := card.Resolve()

	cardJSON, _ := json.Marshal(resolvedCard)
	cardString := string(cardJSON)

	groundTruthString := string(groundTruth)

	if cardString != groundTruthString {
		t.Error(errors.New("ResolvedCard doesn't match the ground truth: " + cardString + " is not " + groundTruthString))
	}
}

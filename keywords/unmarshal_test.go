package keywords

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func allClassesTestCard() CardRootValidator {

	return CardRootValidator{9, &action{
		CardName:       "",
		CastingCost:    3,
		AdditionalCost: nil,
		Class: cardobject.Class{
			Nature:     true,
			Mysticism:  true,
			Technology: true,
			Culture:    true},
		Effects:     nil,
		FlavourText: "",
		Tags:        nil,
		Keywords:    nil}, "ACTION"}
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

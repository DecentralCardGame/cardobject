package cardobject

import "encoding/json"

func NewCardFromJson(jsonString string) (card, error) {
	var card card
	err := json.Unmarshal([]byte(jsonString), &card)
	if err != nil {
		return card, err
	}
	err = card.Validate()
	if err != nil {
		return card, err
	}
	return card, nil
}

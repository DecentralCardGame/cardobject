package cardobject

import (
	"encoding/json"
)

type validateable interface {
	validate() error
}

func NewCardFromJson(jsonString string) (cardInterface, error) {
	var card cardInterface
	err := json.Unmarshal([]byte(jsonString), &card)
	if err != nil {
		return card, err
	}
	err = card.validate()
	if err != nil {
		return card, err
	}
	return card, nil
}

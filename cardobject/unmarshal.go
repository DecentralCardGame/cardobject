package cardobject

import "encoding/json"

func Unmarshal(data []byte) (*Card, error) {
	var card Card
	err := json.Unmarshal(data, &card)
	if err != nil {
		return nil, err
	}
	err = card.ValidateRoot()
	if err != nil {
		return nil, err
	}
	return &card, nil
}

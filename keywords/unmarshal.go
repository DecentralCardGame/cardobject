package keywords

import (
	"encoding/json"
)

/*Unmarshal takes card that contains keywords as json-file and
returns a validated keyworded card
*/
func Unmarshal(data []byte) (KeywordedCard, error) {
	var card card
	err := json.Unmarshal(data, &card)
	if err != nil {
		return nil, err
	}
	err = card.Validate()
	if err != nil {
		return nil, err
	}
	return card, nil
}

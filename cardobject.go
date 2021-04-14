package cardobject

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
)

/*UnmarshalRaw deserializes a raw card-json to a validated cardobject.card.
 */
func UnmarshalRaw(data []byte) (*cardobject.Card, error) {
	card, err := cardobject.Unmarshal(data)
	if err != nil {
		return card, err
	}
	return card, err
}

/*UnmarshalKeyworded deserializes a keyworded card-json to a validated cardobject.card.
 */
func UnmarshalKeyworded(data []byte) (*cardobject.Card, error) {
	keywordedCard, err := keywords.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	resolvedCard := keywordedCard.Resolve()
	err = resolvedCard.Validate()
	if err != nil {
		return &resolvedCard, err
	}
	return &resolvedCard, nil
}

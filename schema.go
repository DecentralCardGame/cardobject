package cardobject

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
)

/*KeywordedSchema provides the instructions on how to build a KeywordedCard via a json-schema
 */
func KeywordedSchema() (string, error) {
	dat, err := keywords.Schema()
	return string(dat), err
}

/*CardSchema provides the instructions on how to build a Card via a json-schema
 */
func CardSchema() (string, error) {
	dat, err := cardobject.Schema()
	return string(dat), err
}

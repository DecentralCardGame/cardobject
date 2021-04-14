package cardobject

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
)

/*KeywordedSchema provides the instructions on how to build a KeywordedCard via a json-schema
 */
func KeywordedSchema() string {
	dat, _ := keywords.Schema()
	return string(dat)
}

/*CardSchema provides the instructions on how to build a Card via a json-schema
 */
func CardSchema() string {
	dat, _ := cardobject.Schema()
	return string(dat)
}

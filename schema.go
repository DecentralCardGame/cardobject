package cardobject

import (
	"github.com/DecentralCardGame/cardobject/keywords"
)

/*KeywordedSchema provides the instructions on how to build a KeywordedCard via a json-schema
 */
func KeywordedSchema() string {
	dat, _ := keywords.Schema()
	return string(dat)
}

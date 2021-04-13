package keywords

import "github.com/DecentralCardGame/cardobject/jsonschema"

/*Schema provides the instructions on how to build a KeywordedCard via a json-schema
 */
func Schema() ([]byte, error) {
	schema, err := jsonschema.Reflect(&card{})
	if err != nil {
		return []byte{}, err
	}
	return schema.MarshalJSONForJS("Card")
}

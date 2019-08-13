package cardobject

import "fmt"
import "github.com/xeipuuv/gojsonschema"
import "encoding/json"

func ProcessCard (cardJson string) string {
	if(validateCard(cardJson)) {
        var card cardWrapper
        //cdc.MustUnmarshalJSON([]byte(cardJson), &card)
        err := json.Unmarshal([]byte(cardJson), &card)
        if err != nil {
            fmt.Println("error:", err)
        }
        bytes, err := json.Marshal(card)
        if err != nil {
            fmt.Println("Can't serialize", card)
            return "Can't serialize"
        }
		return string(bytes)
	} else {
		return ""
	}
}

func validateCard(s string) bool {
    schemaLoader := gojsonschema.NewReferenceLoader("file://C:/Users/marius/Documents/goworkspace/src/github.com/DecentralCardGame/cardobject/schema/cardSchema.json")
    documentLoader := gojsonschema.NewStringLoader(s)

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
        return true
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
        return true//false
    }
}
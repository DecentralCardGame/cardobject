package cardobject

import "fmt"
import "encoding/json"
import "github.com/xeipuuv/gojsonschema"
import amino "github.com/tendermint/go-amino"

func ProcessCard (cardJson string) string {
	cdc := amino.NewCodec()
	registerCodec(cdc)
	var card card
	cdc.MustUnmarshalJSON([]byte(cardJson), &card)

	if(validateCard(cardJson)) {
		return serializeCard(deserializeCard(cardJson))
	} else {
		return ""
	}
}

func serializeCard(c card) string {
	bytes, err := json.Marshal(c)
    if err != nil {
        fmt.Println("Can't serialize", c)
        return "Can't serialize"
    }
    return string(bytes)
}

func deserializeCard(s string) card {
	var card entity
	err := json.Unmarshal([]byte(s), &card)
	if err != nil {
		fmt.Println("error:", err)
	}
	return card
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
        return false
    }
}

func registerCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*card)(nil), nil)
	cdc.RegisterConcrete(action{}, "action", nil)
	cdc.RegisterConcrete(entity{}, "entity", nil)
}
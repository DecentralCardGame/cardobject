package main

import "fmt"
import "github.com/xeipuuv/gojsonschema"

func main() {

    schemaLoader := gojsonschema.NewReferenceLoader("file://C:/Users/marius/Documents/goworkspace/src/github.com/DecentralCardGame/cardobject/schema/cardSchema.json")
    documentLoader := gojsonschema.NewReferenceLoader("file://C:/Users/marius/Documents/goworkspace/src/github.com/DecentralCardGame/cardobject/schema/card.json")

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}
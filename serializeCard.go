package cardobject

import "fmt"
import "os"
//import "io/ioutil"
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
		/*
		// read the file with the schema
		file, err := ioutil.ReadFile("./schema/cardSchema.json")
		if err != nil {
			fmt.Println("\x1b[31;1mreading schema/cardschema.json failed\x1b[0m")
		}
		fmt.Println(string(file))
			*/
    //schemaLoader := gojsonschema.NewStringLoader(string(file))
		schemaLoader := gojsonschema.NewReferenceLoader("file://"+os.Getenv("$GOPATH")+"schema/cardSchema.json")
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

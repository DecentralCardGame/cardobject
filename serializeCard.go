package cardobject

import "path"
import "runtime"
import "github.com/xeipuuv/gojsonschema"
import "encoding/json"
import "errors"

func ProcessCard(cardJson string) (string, error) {
    valid, err := validateCard(cardJson)
	if(valid) {
        var card cardWrapper
        err := json.Unmarshal([]byte(cardJson), &card)
        if err != nil {
            return "Can't deserialize", err
        }
        bytes, err := json.Marshal(card)
        if err != nil {
            return "Can't serialize", err
        }
		return string(bytes), nil
	} else {
		return "", err
	}
}

func validateCard(s string) (bool, error) {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join("file://", path.Dir(filename), "/schema/cardSchema.json")

	schemaLoader := gojsonschema.NewReferenceLoader(filepath)
    documentLoader := gojsonschema.NewStringLoader(s)

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        return true, nil
    } else {
        var errorMessage string
        for _, desc := range result.Errors() {
            errorMessage = errorMessage + ("- %s\n"+ desc.String())
        }
        return false, errors.New(errorMessage)
    }
}

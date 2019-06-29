package cardobject

import "encoding/json"
import "reflect"
import "fmt"
import "errors"

type serializationWrapper struct {
	Type string
}

type cardWrapper struct {
	*serializationWrapper
	Type string
	Value card
}

func (sw *serializationWrapper) Get() string {
	switch wt := reflect.TypeOf(sw); wt.Name() {
		case "reflect.String":
			fmt.Println(wt.Name())
		default:
			fmt.Printf("unhandled kind %s", wt.Name())
		}
	return sw.Type
}

func (cw *cardWrapper) UnmarshalJSON(data []byte) error {
	jsonData := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}
	var cardType string
	err = json.Unmarshal(jsonData["type"], &cardType)
	if err != nil {
		return err
	}
	cw.Type = cardType
	switch cardType {
		case "entity":
			e := entity{}
			err = json.Unmarshal(jsonData["value"], &e)
			if err != nil {
				return err
			}
			cw.Value = e
		case "field":
			f := field{}
			err:= json.Unmarshal(jsonData["value"], &f)
			if err != nil {
				return err
			}
			cw.Value = f
		default:
			return errors.New("Unrecognized card")
	}
	return nil
}

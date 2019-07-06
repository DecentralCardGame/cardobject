package cardobject

import "encoding/json"
import "errors"

type selectorWrapper struct {
	*serializationWrapper
	Type string
	Value selector
}

func (sw *selectorWrapper) UnmarshalJSON(data []byte) error {
	jsonType, jsonData, err := getUnmarshalType(data) 
	sw.Type = jsonType
	switch jsonType {
		case "actionSelector":
			a := actionSelector{}
			err = json.Unmarshal(jsonData, &a)
			sw.Value = a
		case "entitySelector":
			e := entitySelector{}
			err = json.Unmarshal(jsonData, &e)
			sw.Value = e
		case "fieldSelector":
			f := fieldSelector{}
			err = json.Unmarshal(jsonData, &f)
			sw.Value = f
		default:
			return errors.New("Unrecognized selector")
	}
	if err != nil {
		return err
	}
	return nil
}
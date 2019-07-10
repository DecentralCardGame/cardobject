package cardobject

import "encoding/json"
import "errors"

type conditionWrapper struct {
	*serializationWrapper
	Type string
	Value condition
}

func (mw *conditionWrapper) UnmarshalJSON(data []byte) error {
	jsonType, jsonData, err := getUnmarshalType(data) 
	mw.Type = jsonType
	switch jsonType {
		case "intCondition":
			i := intCondition{}
			err = json.Unmarshal(jsonData, &i)
			mw.Value = i
		case "stringCondition":
			s := stringCondition{}
			err = json.Unmarshal(jsonData, &s)
			mw.Value = s
		default:
			return errors.New("Unrecognized condition")
	}
	if err != nil {
		return err
	}
	return nil
}
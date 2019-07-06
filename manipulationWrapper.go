package cardobject

import "encoding/json"
import "errors"

type manipulationWrapper struct {
	*serializationWrapper
	Type string
	Value manipulation
}

func (mw *manipulationWrapper) UnmarshalJSON(data []byte) error {
	jsonType, jsonData, err := getUnmarshalType(data) 
	mw.Type = jsonType
	switch jsonType {
		case "intManipulation":
			i := intManipulation{}
			err = json.Unmarshal(jsonData, &i)
			mw.Value = i
		case "stringManipulation":
			s := stringManipulation{}
			err = json.Unmarshal(jsonData, &s)
			mw.Value = s
		default:
			return errors.New("Unrecognized manipulation")
	}
	if err != nil {
		return err
	}
	return nil
}
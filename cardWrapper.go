package cardobject

import "encoding/json"
import "errors"

type cardWrapper struct {
	*serializationWrapper
	Type string
	Value Card
}

func (cw *cardWrapper) UnmarshalJSON(data []byte) error {
	jsonType, jsonData, err := getUnmarshalType(data) 
	cw.Type = jsonType
	switch jsonType {
		case "action":
			a := action{}
			err = json.Unmarshal(jsonData, &a)
			cw.Value = a
		case "entity":
			e := entity{}
			err = json.Unmarshal(jsonData, &e)
			cw.Value = e
		case "field":
			f := field{}
			err = json.Unmarshal(jsonData, &f)
			cw.Value = f
		case "headquarter":
			h := headquarter{}
			err = json.Unmarshal(jsonData, &h)
			cw.Value = h
		default:
			return errors.New("Unrecognized card")
	}
	if err != nil {
		return err
	}
	return nil
}
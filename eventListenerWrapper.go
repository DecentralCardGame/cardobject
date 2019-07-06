package cardobject

import "encoding/json"
import "errors"

type eventListenerWrapper struct {
	*serializationWrapper
	Type string
	Value eventListener
}

func (ew *eventListenerWrapper) UnmarshalJSON(data []byte) error {
	jsonType, jsonData, err := getUnmarshalType(data) 
	ew.Type = jsonType
	switch jsonType {
		case "timeEventListener":
			t := timeEventListener{}
			err = json.Unmarshal(jsonData, &t)
			ew.Value = t
		case "manipulationEventListener":
			m := manipulationEventListener{}
			err = json.Unmarshal(jsonData, &m)
			ew.Value = m
		case "zoneChangeEventListener":
			z := zoneChangeEventListener{}
			err = json.Unmarshal(jsonData, &z)
			ew.Value = z
		default:
			return errors.New("Unrecognized eventListener")
	}
	if err != nil {
		return err
	}
	return nil
}
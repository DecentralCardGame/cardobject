package cardobject

import "encoding/json"
import "errors"

type abilityWrapper struct {
	*serializationWrapper
	Type string
	Value ability
}

func (aw *abilityWrapper) UnmarshalJSON(data []byte) error {
	jsonType, jsonData, err := getUnmarshalType(data) 
	aw.Type = jsonType
	switch jsonType {
		case "activatedAbility":
			a := activatedAbility{}
			err = json.Unmarshal(jsonData, &a)
			aw.Value = a
		case "triggeredAbility":
			t := triggeredAbility{}
			err = json.Unmarshal(jsonData, &t)
			aw.Value = t
		default:
			return errors.New("Unrecognized ability")
	}
	if err != nil {
		return err
	}
	return nil
}
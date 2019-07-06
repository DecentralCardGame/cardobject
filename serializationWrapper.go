package cardobject

import "encoding/json"

type serializationWrapper struct {
	Type string
}

func getUnmarshalType(data []byte) (string, []byte, error) {
	jsonData := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", nil, err
	}
	var cardType string
	err = json.Unmarshal(jsonData["Type"], &cardType)
	if err != nil {
		return "", nil, err
	}
	return cardType, jsonData["Value"], nil
}

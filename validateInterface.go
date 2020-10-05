package cardobject

import (
	"errors"
	"reflect"
)

func xorInterface(possibleImplementer []validateable) (validateable, error) {
	var implementer validateable
	implementerFound := false
	for _, b := range possibleImplementer {
		if !reflect.ValueOf(b).IsNil() {
			if implementerFound {
				return nil, errors.New("ImplementerError")
			}
			implementer = b
			implementerFound = true
		}
	}
	return implementer, nil
}

func combineErrors(errorRange []error) error {
	isError := false
	errorString := ""
	for _, error := range errorRange {
		if error != nil {
			isError = true
			errorString = errorString + "\n" + error.Error()
		}
	}
	if isError {
		return errors.New(errorString)
	}
	return nil
}

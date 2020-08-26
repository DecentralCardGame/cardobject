package blockchainCard

import "errors"

func xorInterface(possibleImplementer []validateable) validateable {
	var implmenter validateable
	for _, b := range possibleImplementer {
		if b != nil {
			implmenter = b
		}
	}
	return implmenter
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
	} else {
		return nil
	}
}

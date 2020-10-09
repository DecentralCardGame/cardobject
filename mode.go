package cardobject

import "errors"

var intChangeModes []string = []string{"INCREASES", "DECREASES", "CHANGES"}
var stringChangeModes []string = []string{"CHANGES"}

func validateIntChangeMode(mode string) error {
	for _, t := range intChangeModes {
		if t == mode {
			return nil
		}
	}
	return errors.New("IntChangeMode: " + mode + " not available")
}

func validateStringChangeMode(mode string) error {
	for _, t := range stringChangeModes {
		if t == mode {
			return nil
		}
	}
	return errors.New("stringChangeMode: " + mode + " not available")
}

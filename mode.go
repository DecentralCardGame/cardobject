package cardobject

import "errors"

var intChangeModes []string = []string{"INCREASES", "DECREASES", "CHANGES"}
var stringChangeModes []string = []string{"CHANGES"}

var playerModes []string = []string{"YOU", "OPPONENT"}
var cardModes []string = []string{"ALL", "OPPONENTSCHOICE", "RANDOM", "TARGET"}
var ownerModes []string = []string{"YOUR", "OPPONENTS", "OWNERS"}

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

func validatePlayerMode(mode string) error {
	for _, t := range playerModes {
		if t == mode {
			return nil
		}
	}
	return errors.New("playerMode: " + mode + " not available")
}

func validateCardMode(mode string) error {
	for _, t := range cardModes {
		if t == mode {
			return nil
		}
	}
	return errors.New("cardMode: " + mode + " not available")
}
func validateOwnerMode(mode string) error {
	for _, t := range ownerModes {
		if t == mode {
			return nil
		}
	}
	return errors.New("OwnerMode: " + mode + " not available")
}

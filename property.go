package cardobject

import "errors"

var actionIntProperties []string = []string{"COSTSUM"}
var actionStringProperties []string = []string{"NAME", "TEXT"}
var entityIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH"}
var entityStringProperties []string = []string{"NAME", "TEXT"}
var placeIntProperties []string = []string{"COSTSUM", "HEALTH"}
var placeStringProperties []string = []string{"NAME", "TEXT"}

var cardIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH"}
var cardStringProperties []string = []string{"NAME", "TEXT"}

func validateActionIntProperty(property string) error {
	for _, t := range actionIntProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("ActionIntProperty: " + property + " not available")
}

func validateActionStringProperty(property string) error {
	for _, t := range actionStringProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("ActionStringProperty: " + property + " not available")
}

func validateEntityIntProperty(property string) error {
	for _, t := range entityIntProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("EntityIntProperty: " + property + " not available")
}

func validateEntityStringProperty(property string) error {
	for _, t := range entityStringProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("EntityStringProperty: " + property + " not available")
}

func validatePlaceIntProperty(property string) error {
	for _, t := range placeIntProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("PlaceIntProperty: " + property + " not available")
}

func validatePlaceStringProperty(property string) error {
	for _, t := range placeStringProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("PlaceStringProperty: " + property + " not available")
}

func validateCardIntProperty(property string) error {
	for _, t := range cardIntProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("CardIntProperty: " + property + " not available")
}

func validateCardStringProperty(property string) error {
	for _, t := range cardStringProperties {
		if t == property {
			return nil
		}
	}
	return errors.New("CardStringProperty: " + property + " not available")
}

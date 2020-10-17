package cardobject

import "errors"

var dynamicZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND"}
var zones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

var actionZones []string = []string{"DECK", "DUSTPILE", "HAND", "VOID"}
var entityZones []string = []string{"ATTACKLANE", "BLOCKLANE", "DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}
var placeZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

func validateDynamicZone(zone string) error {
	for _, t := range dynamicZones {
		if t == zone {
			return nil
		}
	}
	return errors.New("DynamicZone: " + zone + " not available")
}

func validateZone(zone string) error {
	for _, t := range zones {
		if t == zone {
			return nil
		}
	}
	return errors.New("Zone: " + zone + " not available")
}

func validateActionZone(zone string) error {
	for _, t := range actionZones {
		if t == zone {
			return nil
		}
	}
	return errors.New("ActionZone: " + zone + " not available")
}

func validateEntityZone(zone string) error {
	for _, t := range entityZones {
		if t == zone {
			return nil
		}
	}
	return errors.New("EntityZone: " + zone + " not available")
}
func validatePlaceZone(zone string) error {
	for _, t := range placeZones {
		if t == zone {
			return nil
		}
	}
	return errors.New("PlaceZone: " + zone + " not available")
}

package cardobject

import "errors"

var dynamicZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND"}
var zones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

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

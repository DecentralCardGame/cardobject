package cardobject

import "errors"

var timeEvents []string = []string{"TICKSTART", "COMBAT"}

func validateTimeEvent(timeEvent string) error {
	for _, t := range timeEvents {
		if t == timeEvent {
			return nil
		}
	}
	return errors.New("TimeEvent: " + timeEvent + " not available")
}

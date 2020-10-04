package cardobject

import "errors"

var intComparators []string = []string{"EQUAL", "GREATER", "LESSER"}
var stringComparators []string = []string{"EQUAL", "CONTAINS", "UNEQUAL", "CONTAINSNOT"}

func validateIntComparator(comparator string) error {
	for _, t := range intComparators {
		if t == comparator {
			return nil
		}
	}
	return errors.New("IntComparator: " + comparator + " not available")
}

func validateStringComparator(comparator string) error {
	for _, t := range stringComparators {
		if t == comparator {
			return nil
		}
	}
	return errors.New("StringComparator: " + comparator + " not available")
}

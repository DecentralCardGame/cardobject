package blockchainCard

import (
	"errors"
	"strconv"
)

var possibleTags []string = []string{
	"ANIMAL",
	"BOT",
	"DWARF",
	"ENGINEER",
	"EQUIPMENT",
	"FARM",
	"FIRE",
	"HUMAN",
	"KNIGHT",
	"MAGIC",
	"MILITANT",
	"PRIMITIVE",
	"RANGE",
	"SPIRITUAL",
	"TACTIC",
	"TECHNOCRAT"}

func validateTags(tags []string) error {
	if len(tags) > maxTagCount || len(tags) < minTagCount {
		return errors.New("The card must have at least " + strconv.Itoa(minTagCount) + " and at most " + strconv.Itoa(maxTagCount) + "tags")
	}
	errorRange := []error{}
	for _, tag := range tags {
		errorRange = append(errorRange, validateTag(tag))
	}
	return combineErrors(errorRange)
}

func validateTag(tag string) error {
	for _, t := range possibleTags {
		if t == tag {
			return nil
		}
	}
	return errors.New("Tag not available")
}

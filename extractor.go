package cardobject

import (
	"errors"
	"strconv"
)

type cardExtractors struct {
	ActionExtractors *actionExtractors `json:",omitempty"`
	EntityExtractors *entityExtractors `json:",omitempty"`
	PlaceExtractors  *placeExtractors  `json:",omitempty"`
}

type actionExtractors []actionExtractor

type entityExtractors []entityExtractor

type placeExtractors []placeExtractor

type actionExtractor struct {
	ActionIntExtractor    *actionIntExtractor    `json:",omitempty"`
	ActionStringExtractor *actionStringExtractor `json:",omitempty"`
	ActionTargetExtractor *targetExtractor       `json:",omitempty"`
}

type entityExtractor struct {
	EntityIntExtractor    *entityIntExtractor    `json:",omitempty"`
	EntityStringExtractor *entityStringExtractor `json:",omitempty"`
	EntityTargetExtractor *targetExtractor       `json:",omitempty"`
}

type placeExtractor struct {
	PlaceIntExtractor    *placeIntExtractor    `json:",omitempty"`
	PlaceStringExtractor *placeStringExtractor `json:",omitempty"`
	PlaceTargetExtractor *targetExtractor      `json:",omitempty"`
}

type actionIntExtractor struct {
	ExtractIntProperty string
	IntVariableName    string
}
type actionStringExtractor struct {
	ExtractStringProperty string
	StringVariableName    string
}

type entityIntExtractor struct {
	ExtractIntProperty string
	IntVariableName    string
}
type entityStringExtractor struct {
	ExtractStringProperty string
	StringVariableName    string
}

type placeIntExtractor struct {
	ExtractIntProperty string
	IntVariableName    string
}
type placeStringExtractor struct {
	ExtractStringProperty string
	StringVariableName    string
}

type targetExtractor struct {
	TargetVariableName string
}

type intExtractor struct {
	IntVariableName string
}

type stringExtractor struct {
	StringVariableName string
}

func (e *cardExtractors) validate() error {
	possibleImplementer := []validateable{e.ActionExtractors, e.EntityExtractors, e.PlaceExtractors}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Extractor implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e actionExtractors) validate() error {
	if len(e) > maxExtractorCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxExtractorCount) + " extractors")
	}
	errorRange := []error{}
	for _, actionExtractor := range e {
		errorRange = append(errorRange, actionExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (e entityExtractors) validate() error {
	if len(e) > maxExtractorCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxExtractorCount) + " extractors")
	}
	errorRange := []error{}
	for _, entityExtractor := range e {
		errorRange = append(errorRange, entityExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (e placeExtractors) validate() error {
	if len(e) > maxExtractorCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxExtractorCount) + " extractors")
	}
	errorRange := []error{}
	for _, placeExtractor := range e {
		errorRange = append(errorRange, placeExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (e *actionExtractor) validate() error {
	possibleImplementer := []validateable{e.ActionIntExtractor, e.ActionStringExtractor, e.ActionTargetExtractor}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("ActionExtractor implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *entityExtractor) validate() error {
	possibleImplementer := []validateable{e.EntityIntExtractor, e.EntityStringExtractor, e.EntityTargetExtractor}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("ActionExtractor implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *placeExtractor) validate() error {
	possibleImplementer := []validateable{e.PlaceIntExtractor, e.PlaceStringExtractor, e.PlaceTargetExtractor}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("ActionExtractor implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *actionIntExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionIntProperty(e.ExtractIntProperty))
	errorRange = append(errorRange, validateIntVariableName(e.IntVariableName))
	return combineErrors(errorRange)
}

func (e *actionStringExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionIntProperty(e.ExtractStringProperty))
	errorRange = append(errorRange, validateStringVariableName(e.StringVariableName))
	return combineErrors(errorRange)
}

func (e *entityIntExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityIntProperty(e.ExtractIntProperty))
	errorRange = append(errorRange, validateIntVariableName(e.IntVariableName))
	return combineErrors(errorRange)
}

func (e *entityStringExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityIntProperty(e.ExtractStringProperty))
	errorRange = append(errorRange, validateStringVariableName(e.StringVariableName))
	return combineErrors(errorRange)
}

func (e *placeIntExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceIntProperty(e.ExtractIntProperty))
	errorRange = append(errorRange, validateIntVariableName(e.IntVariableName))
	return combineErrors(errorRange)
}

func (e *placeStringExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceIntProperty(e.ExtractStringProperty))
	errorRange = append(errorRange, validateStringVariableName(e.StringVariableName))
	return combineErrors(errorRange)
}

func (e *targetExtractor) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringVariableName(e.TargetVariableName))
	return combineErrors(errorRange)
}

func (e *intExtractor) validate() error {
	return validateIntVariableName(e.IntVariableName)
}

func (e *stringExtractor) validate() error {
	return validateStringVariableName(e.StringVariableName)
}

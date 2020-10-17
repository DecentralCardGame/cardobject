package cardobject

import (
	"errors"
	"strconv"
)

const maxAttack int = 32
const minAttack int = 0
const maxCardNameLength int = 32
const minCardNameLength int = 1
const maxCastingCost int = 32
const minCastingCost int = 0
const maxConditionCount int = 3
const maxExtractorCount int = 3
const maxAbilityEffectCount int = 3
const maxManipulationCount int = 3
const maxFlavourTextLength int = 1000
const minFlavourTextLength int = 1
const maxGrowth = 32
const minGrowth = 0
const maxHealth int = 32
const minHealth int = 0
const maxSimpleInt int = 32
const minSimpleInt int = 0
const maxSimpleStringLength = 32
const minSimpleStringLength = 1
const maxStartingHandSize int = 10
const minStartingHandSize int = 0
const maxTagCount int = 3
const minTagCount int = 1
const maxWisdom int = 32
const minWisdom int = 0

func validateAttack(attack int) error {
	if attack > maxAttack || attack < minAttack {
		return errors.New("The cards attack must be between " + strconv.Itoa(minAttack) + " and " + strconv.Itoa(maxAttack))
	}
	return nil
}

func validateCardName(cardName string) error {
	cardNameLength := len(cardName)
	if cardNameLength > maxCardNameLength || cardNameLength < minCardNameLength {
		return errors.New("The cards name must have between " + strconv.Itoa(minCardNameLength) + " and " + strconv.Itoa(maxCardNameLength) + " characters")
	}
	return nil
}

func validateCastingCost(castingCost int) error {
	if castingCost > maxCastingCost || castingCost < 0 {
		return errors.New("The casting must be between " + strconv.Itoa(minCastingCost) + " and " + strconv.Itoa(maxCastingCost))
	}
	return nil
}

func validateFlavourText(flavourText string) error {
	flavourTextLength := len(flavourText)
	if flavourTextLength > maxFlavourTextLength || flavourTextLength < minFlavourTextLength {
		return errors.New("The cards flavour text must have between " + strconv.Itoa(minFlavourTextLength) + " and " + strconv.Itoa(maxFlavourTextLength) + " characters")
	}
	return nil
}

func validateGrowth(growth int) error {
	if growth > maxGrowth || growth < minGrowth {
		return errors.New("The cards growth must be between " + strconv.Itoa(minGrowth) + " and " + strconv.Itoa(maxGrowth))
	}
	return nil
}

func validateHealth(health int) error {
	if health > maxHealth || health < minHealth {
		return errors.New("The cards health must be between " + strconv.Itoa(minHealth) + " and " + strconv.Itoa(maxHealth))
	}
	return nil
}

func validateStartingHandSize(startingHandSize int) error {
	if startingHandSize > maxStartingHandSize || startingHandSize < minStartingHandSize {
		return errors.New("The cards startingHandSize must be between " + strconv.Itoa(minStartingHandSize) + " and " + strconv.Itoa(minStartingHandSize))
	}
	return nil
}

func validateWisdom(wisdom int) error {
	if wisdom > maxWisdom || wisdom < minWisdom {
		return errors.New("The cards growth must be between " + strconv.Itoa(minWisdom) + " and " + strconv.Itoa(maxWisdom))
	}
	return nil
}

func validateSimpleInt(simpleInt int) error {
	if simpleInt > maxSimpleInt || simpleInt < minSimpleInt {
		return errors.New("Simple Integer must be between " + strconv.Itoa(minSimpleInt) + " and " + strconv.Itoa(maxSimpleInt))
	}
	return nil
}

func validateSimpleString(simpleString string) error {
	simpleStringLength := len(simpleString)
	if simpleStringLength > maxSimpleStringLength || simpleStringLength < minSimpleStringLength {
		return errors.New("Simple String-length mus be between " + strconv.Itoa(minSimpleStringLength) + " and " + strconv.Itoa(maxSimpleStringLength))
	}
	return nil
}

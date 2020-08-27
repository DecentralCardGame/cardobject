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
const maxAbilityEffectCount int = 3
const maxFlavourTextLength int = 1000
const minFlavourTextLength int = 1
const maxGrowth = 32
const minGrowth = 0
const maxHealth int = 32
const minHealth int = 0
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
	if len(cardName) > maxCardNameLength {
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
	if len(flavourText) > maxFlavourTextLength {
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

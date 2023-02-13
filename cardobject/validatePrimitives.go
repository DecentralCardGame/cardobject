package cardobject

import (
	"errors"
	"strconv"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

const maxAttack int = 32
const minAttack int = 0
const maxBasicAmount = 32
const minBasicAmount = 0
const maxCardNameLength int = 25
const minCardNameLength int = 1
const maxCastingCost int = 64
const minCastingCost int = 0
const maxConditionCount int = 3
const maxDelay int = 32
const minDelay int = 1
const maxExtractorCount int = 3
const maxAbilityEffectCount int = 3
const maxManipulationCount int = 3
const maxFlavourTextLength int = 1000
const minFlavourTextLength int = 0
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

type Attack jsonschema.BasicInt

func (a Attack) ValidateType(r jsonschema.RootElement) error {
	min, max := a.MinMax()
	if a < Attack(min) || a > Attack(max) {
		return errors.New("Attack must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (a Attack) MinMax() (int, int) {
	return minAttack, maxAttack
}

type BasicAmount jsonschema.BasicInt

func (b BasicAmount) ValidateType(r jsonschema.RootElement) error {
	min, max := b.MinMax()
	if b < BasicAmount(min) || b > BasicAmount(max) {
		return errors.New("BasicAmount must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (b BasicAmount) MinMax() (int, int) {
	return minBasicAmount, maxBasicAmount
}

type CardName jsonschema.BasicString

func (c CardName) ValidateType(r jsonschema.RootElement) error {
	minLength, maxLength := c.MinMaxLength()
	length := len(string(c))
	if length < minLength || length > maxLength {
		return errors.New("CardName must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (c CardName) MinMaxLength() (int, int) {
	return minCardNameLength, maxCardNameLength
}

type CastingCost jsonschema.BasicInt

func (c CastingCost) ValidateType(r jsonschema.RootElement) error {
	min, max := c.MinMax()
	if c < CastingCost(min) || c > CastingCost(max) {
		return errors.New("CastingCost must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (c CastingCost) MinMax() (int, int) {
	return minCastingCost, maxCastingCost
}

type Delay jsonschema.BasicInt

func (d Delay) ValidateType(r jsonschema.RootElement) error {
	min, max := d.MinMax()
	if d < Delay(min) || d > Delay(max) {
		return errors.New("Delay must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (d Delay) MinMax() (int, int) {
	return minDelay, maxDelay
}

type FlavourText jsonschema.BasicString

func (f FlavourText) ValidateType(r jsonschema.RootElement) error {
	minLength, maxLength := f.MinMaxLength()
	length := len(string(f))
	if length < minLength || length > maxLength {
		return errors.New("FlavourText must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (f FlavourText) MinMaxLength() (int, int) {
	return minFlavourTextLength, maxFlavourTextLength
}

type Growth jsonschema.BasicInt

func (g Growth) ValidateType(r jsonschema.RootElement) error {
	min, max := g.MinMax()
	if g < Growth(min) || g > Growth(max) {
		return errors.New("Growth must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (g Growth) MinMax() (int, int) {
	return minGrowth, maxGrowth
}

type Health jsonschema.BasicInt

func (h Health) ValidateType(r jsonschema.RootElement) error {
	min, max := h.MinMax()
	if h < Health(min) || h > Health(max) {
		return errors.New("Health must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (h Health) MinMax() (int, int) {
	return minHealth, maxHealth
}

type StartingHandsize jsonschema.BasicInt

func (s StartingHandsize) ValidateType(r jsonschema.RootElement) error {
	min, max := s.MinMax()
	if s < StartingHandsize(min) || s > StartingHandsize(max) {
		return errors.New("StartingHandsize must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (s StartingHandsize) MinMax() (int, int) {
	return minStartingHandSize, maxStartingHandSize
}

type Wisdom jsonschema.BasicInt

func (w Wisdom) ValidateType(r jsonschema.RootElement) error {
	min, max := w.MinMax()
	if w < Wisdom(min) || w > Wisdom(max) {
		return errors.New("Wisdom must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (w Wisdom) MinMax() (int, int) {
	return minWisdom, maxWisdom
}

type SimpleIntValue jsonschema.BasicInt

func (s SimpleIntValue) ValidateType(r jsonschema.RootElement) error {
	min, max := s.MinMax()
	if s < SimpleIntValue(min) || s > SimpleIntValue(max) {
		return errors.New("SimpleIntValue must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (s SimpleIntValue) MinMax() (int, int) {
	return minSimpleInt, maxSimpleInt
}

type SimpleStringValue jsonschema.BasicString

func (s SimpleStringValue) ValidateType(r jsonschema.RootElement) error {
	minLength, maxLength := s.MinMaxLength()
	length := len(string(s))
	if length < minLength || length > maxLength {
		return errors.New("CardName must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (s SimpleStringValue) MinMaxLength() (int, int) {
	return minSimpleStringLength, maxSimpleStringLength
}

package cardobject

import (
	"errors"
	"strconv"

	"github.com/DecentralCardGame/jsonschema"
)

const maxAttack int = 32
const minAttack int = 0
const maxBasicAmount = 32
const minBasicAmount = 0
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

type Attack jsonschema.BasicInt

func (a Attack) Validate() error {
	return a.ValidateInt()
}

func (a Attack) ValidateInt() error {
	min, max := a.GetMinMax()
	if a < Attack(min) || a > Attack(max) {
		return errors.New("Attack must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (a Attack) GetMinMax() (int, int) {
	return minAttack, maxAttack
}

type BasicAmount jsonschema.BasicInt

func (b BasicAmount) Validate() error {
	return b.ValidateInt()
}

func (b BasicAmount) ValidateInt() error {
	min, max := b.GetMinMax()
	if b < BasicAmount(min) || b > BasicAmount(max) {
		return errors.New("BasicAmount must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (b BasicAmount) GetMinMax() (int, int) {
	return minBasicAmount, maxBasicAmount
}

type CardName jsonschema.BasicString

func (c CardName) Validate() error {
	return c.ValidateString()
}

func (c CardName) ValidateString() error {
	minLength, maxLength := c.GetMinMaxLength()
	length := len(string(c))
	if length < minLength || length > maxLength {
		return errors.New("CardName must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (c CardName) GetMinMaxLength() (int, int) {
	return minCardNameLength, maxCardNameLength
}

type CastingCost jsonschema.BasicInt

func (c CastingCost) Validate() error {
	return c.ValidateInt()
}

func (c CastingCost) ValidateInt() error {
	min, max := c.GetMinMax()
	if c < CastingCost(min) || c > CastingCost(max) {
		return errors.New("CastingCost must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (c CastingCost) GetMinMax() (int, int) {
	return minCastingCost, maxCastingCost
}

type FlavourText jsonschema.BasicString

func (f FlavourText) Validate() error {
	return f.ValidateString()
}

func (f FlavourText) ValidateString() error {
	minLength, maxLength := f.GetMinMaxLength()
	length := len(string(f))
	if length < minLength || length > maxLength {
		return errors.New("FlavourText must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (f FlavourText) GetMinMaxLength() (int, int) {
	return minFlavourTextLength, maxFlavourTextLength
}

type Growth jsonschema.BasicInt

func (g Growth) Validate() error {
	return g.ValidateInt()
}

func (g Growth) ValidateInt() error {
	min, max := g.GetMinMax()
	if g < Growth(min) || g > Growth(max) {
		return errors.New("Growth must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (g Growth) GetMinMax() (int, int) {
	return minGrowth, maxGrowth
}

type Health jsonschema.BasicInt

func (h Health) Validate() error {
	return h.ValidateInt()
}

func (h Health) ValidateInt() error {
	min, max := h.GetMinMax()
	if h < Health(min) || h > Health(max) {
		return errors.New("Health must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (h Health) GetMinMax() (int, int) {
	return minHealth, maxHealth
}

type StartingHandsize jsonschema.BasicInt

func (s StartingHandsize) Validate() error {
	return s.ValidateInt()
}

func (s StartingHandsize) ValidateInt() error {
	min, max := s.GetMinMax()
	if s < StartingHandsize(min) || s > StartingHandsize(max) {
		return errors.New("StartingHandsize must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (s StartingHandsize) GetMinMax() (int, int) {
	return minStartingHandSize, maxStartingHandSize
}

type Wisdom jsonschema.BasicInt

func (w Wisdom) Validate() error {
	return w.ValidateInt()
}

func (w Wisdom) ValidateInt() error {
	min, max := w.GetMinMax()
	if w < Wisdom(min) || w > Wisdom(max) {
		return errors.New("Wisdom must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (w Wisdom) GetMinMax() (int, int) {
	return minWisdom, maxWisdom
}

type SimpleIntValue jsonschema.BasicInt

func (s SimpleIntValue) Validate() error {
	return s.ValidateInt()
}

func (s SimpleIntValue) ValidateInt() error {
	min, max := s.GetMinMax()
	if s < SimpleIntValue(min) || s > SimpleIntValue(max) {
		return errors.New("SimpleIntValue must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func (s SimpleIntValue) GetMinMax() (int, int) {
	return minSimpleInt, maxSimpleInt
}

type SimpleStringValue jsonschema.BasicString

func (s SimpleStringValue) Validate() error {
	return s.ValidateString()
}

func (s SimpleStringValue) ValidateString() error {
	minLength, maxLength := s.GetMinMaxLength()
	length := len(string(s))
	if length < minLength || length > maxLength {
		return errors.New("CardName must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func (s SimpleStringValue) GetMinMaxLength() (int, int) {
	return minSimpleStringLength, maxSimpleStringLength
}

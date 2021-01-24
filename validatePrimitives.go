package cardobject

import (
	"errors"

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

type attack jsonschema.BasicInt

func (a attack) Validate() error {
	return a.ValidateInt()
}

func (a attack) ValidateInt() error {
	min, max := a.GetMinMax()
	if a < attack(min) || a > attack(max) {
		return errors.New("")
	}
	return nil
}

func (a attack) GetMinMax() (int, int) {
	return minAttack, maxAttack
}

type basicAmount jsonschema.BasicInt

func (b basicAmount) Validate() error {
	return b.ValidateInt()
}

func (b basicAmount) ValidateInt() error {
	min, max := b.GetMinMax()
	if b < basicAmount(min) || b > basicAmount(max) {
		return errors.New("")
	}
	return nil
}

func (b basicAmount) GetMinMax() (int, int) {
	return minBasicAmount, maxBasicAmount
}

type cardName jsonschema.BasicString

func (c cardName) Validate() error {
	return c.ValidateString()
}

func (c cardName) ValidateString() error {
	minLength, maxLength := c.GetMinMaxLength()
	length := len(string(c))
	if length < minLength || length > maxLength {
		return errors.New("")
	}
	return nil
}

func (c cardName) GetMinMaxLength() (int, int) {
	return minCardNameLength, maxCardNameLength
}

type castingCost jsonschema.BasicInt

func (c castingCost) Validate() error {
	return c.ValidateInt()
}

func (c castingCost) ValidateInt() error {
	min, max := c.GetMinMax()
	if c < castingCost(min) || c > castingCost(max) {
		return errors.New("")
	}
	return nil
}

func (c castingCost) GetMinMax() (int, int) {
	return minCastingCost, maxCastingCost
}

type flavourText jsonschema.BasicString

func (f flavourText) Validate() error {
	return f.ValidateString()
}

func (f flavourText) ValidateString() error {
	minLength, maxLength := f.GetMinMaxLength()
	length := len(string(f))
	if length < minLength || length > maxLength {
		return errors.New("")
	}
	return nil
}

func (f flavourText) GetMinMaxLength() (int, int) {
	return minFlavourTextLength, maxFlavourTextLength
}

type growth jsonschema.BasicInt

func (g growth) Validate() error {
	return g.ValidateInt()
}

func (g growth) ValidateInt() error {
	min, max := g.GetMinMax()
	if g < growth(min) || g > growth(max) {
		return errors.New("")
	}
	return nil
}

func (g growth) GetMinMax() (int, int) {
	return minGrowth, maxGrowth
}

type health jsonschema.BasicInt

func (h health) Validate() error {
	return h.ValidateInt()
}

func (h health) ValidateInt() error {
	min, max := h.GetMinMax()
	if h < health(min) || h > health(max) {
		return errors.New("")
	}
	return nil
}

func (h health) GetMinMax() (int, int) {
	return minHealth, maxHealth
}

type startingHandsize jsonschema.BasicInt

func (s startingHandsize) Validate() error {
	return s.ValidateInt()
}

func (s startingHandsize) ValidateInt() error {
	min, max := s.GetMinMax()
	if s < startingHandsize(min) || s > startingHandsize(max) {
		return errors.New("")
	}
	return nil
}

func (s startingHandsize) GetMinMax() (int, int) {
	return minStartingHandSize, maxStartingHandSize
}

type wisdom jsonschema.BasicInt

func (w wisdom) Validate() error {
	return w.ValidateInt()
}

func (w wisdom) ValidateInt() error {
	min, max := w.GetMinMax()
	if w < wisdom(min) || w > wisdom(max) {
		return errors.New("")
	}
	return nil
}

func (w wisdom) GetMinMax() (int, int) {
	return minWisdom, maxWisdom
}

type simpleIntValue jsonschema.BasicInt

func (s simpleIntValue) Validate() error {
	return s.ValidateInt()
}

func (s simpleIntValue) ValidateInt() error {
	min, max := s.GetMinMax()
	if s < simpleIntValue(min) || s > simpleIntValue(max) {
		return errors.New("")
	}
	return nil
}

func (s simpleIntValue) GetMinMax() (int, int) {
	return minSimpleInt, maxSimpleInt
}

type simpleStringValue jsonschema.BasicString

func (s simpleStringValue) Validate() error {
	return s.ValidateString()
}

func (s simpleStringValue) ValidateString() error {
	minLength, maxLength := s.GetMinMaxLength()
	length := len(string(s))
	if length < minLength || length > maxLength {
		return errors.New("")
	}
	return nil
}

func (s simpleStringValue) GetMinMaxLength() (int, int) {
	return minSimpleStringLength, maxSimpleStringLength
}

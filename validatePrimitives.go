package cardobject

import (
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

type attack struct{ *jsonschema.BasicInt }

func (a attack) GetMinMax() (int, int) {
	return minAttack, maxAttack
}

type basicAmount struct{ *jsonschema.BasicInt }

func (b basicAmount) GetMinMax() (int, int) {
	return minBasicAmount, maxBasicAmount
}

type cardName struct{ *jsonschema.BasicString }

func (c cardName) GetMinMaxLength() (int, int) {
	return minCardNameLength, maxCardNameLength
}

type castingCost struct{ *jsonschema.BasicInt }

func (c castingCost) GetMinMax() (int, int) {
	return minCastingCost, maxCastingCost
}

type flavourText struct{ *jsonschema.BasicString }

func (f flavourText) GetMinMaxLength() (int, int) {
	return minFlavourTextLength, maxFlavourTextLength
}

type growth struct{ *jsonschema.BasicInt }

func (h growth) GetMinMax() (int, int) {
	return minGrowth, maxGrowth
}

type health struct{ *jsonschema.BasicInt }

func (h health) GetMinMax() (int, int) {
	return minHealth, maxHealth
}

type startingHandsize struct{ *jsonschema.BasicInt }

func (s startingHandsize) GetMinMax() (int, int) {
	return minStartingHandSize, maxStartingHandSize
}

type wisdom struct{ *jsonschema.BasicInt }

func (w wisdom) GetMinMax() (int, int) {
	return minWisdom, maxWisdom
}

type simpleIntValue struct{ *jsonschema.BasicInt }

func (s simpleIntValue) GetMinMax() (int, int) {
	return minSimpleInt, maxSimpleInt
}

type simpleStringValue struct{ *jsonschema.BasicString }

func (s simpleStringValue) GetMinMaxLength() (int, int) {
	return minSimpleStringLength, maxSimpleStringLength
}

package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//XVar VariableName
const XVar = "X"

//YVar VariableName
const YVar = "Y"

//ZVar VariableName
const ZVar = "Z"

//AVar VariableName
const AVar = "A"

//BVar VariableName
const BVar = "B"

//CVar VariableName
const CVar = "C"

//MVar VariableName
const MVar = "M"

//TVar VariableName
const TVar = "T"

var intVariableNames []string = []string{XVar, YVar, ZVar}
var stringVariableNames []string = []string{AVar, BVar, CVar}
var targetVariableNames []string = []string{MVar, TVar}

type IntVariableName jsonschema.BasicEnum

func (i IntVariableName) Validate() error {
	return i.ValidateEnum()
}

func (i IntVariableName) ValidateEnum() error {
	values := i.EnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("IntVariableName must be on of " + strings.Join(intVariableNames, ","))
}

func (i IntVariableName) EnumValues() []string {
	return intVariableNames
}

type StringVariableName jsonschema.BasicEnum

func (s StringVariableName) Validate() error {
	return s.ValidateEnum()
}

func (s StringVariableName) ValidateEnum() error {
	values := s.EnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("StringVariableName must be on of " + strings.Join(stringVariableNames, ","))
}

func (s StringVariableName) EnumValues() []string {
	return stringVariableNames
}

type TargetVariableName jsonschema.BasicEnum

func (t TargetVariableName) Validate() error {
	return t.ValidateEnum()
}

func (t TargetVariableName) ValidateEnum() error {
	values := t.EnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("TargetVariableName must be on of " + strings.Join(targetVariableNames, ","))
}

func (t TargetVariableName) EnumValues() []string {
	return targetVariableNames
}

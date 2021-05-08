package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//Recruit Token
const Recruit = "1/1 human Recruit"

//Bot Token
const Bot = "2/2 Bot"

//Beast Token
const Beast = "3/3 animal Beast"

var tokenTypes []string = []string{Recruit, Bot, Beast}

type TokenType jsonschema.BasicEnum

func (t TokenType) Validate() error {
	return t.ValidateEnum()
}

func (t TokenType) ValidateEnum() error {
	values := t.EnumValues()
	for _, t := range values {
		if t == string(t) {
			return nil
		}
	}
	return errors.New("TokenType must be one of: " + strings.Join(tokenTypes, ","))
}

func (t TokenType) EnumValues() []string {
	return tokenTypes
}

type Token struct {
	Name   CardName `json:",omitempty"`
	Attack Attack
	Health Health
	Tags   Tags `json:",omitempty"`
}

func (t Token) Validate() error {
	return t.ValidateStruct()
}

func (t Token) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t Token) InteractionText() string {
	return "§Attack / §Health token named §Name tagged §Tags"
}

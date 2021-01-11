package cardobject

import "github.com/DecentralCardGame/jsonschema"

var intChangeModes []string = []string{"INCREASES", "DECREASES", "CHANGES"}
var stringChangeModes []string = []string{"CHANGES"}

var playerModes []string = []string{"YOU", "OPPONENT"}
var cardModes []string = []string{"ALL", "OPPONENTSCHOICE", "RANDOM", "TARGET"}
var ownerModes []string = []string{"YOUR", "OPPONENTS", "OWNERS"}

type intChangeMode struct{ *jsonschema.BasicEnum }

func (i intChangeMode) GetEnumValues() []string {
	return intChangeModes
}

type stringChangeMode struct{ *jsonschema.BasicEnum }

func (s stringChangeMode) GetEnumValues() []string {
	return stringChangeModes
}

type playerMode struct{ *jsonschema.BasicEnum }

func (p playerMode) GetEnumValues() []string {
	return playerModes
}

type cardMode struct{ *jsonschema.BasicEnum }

func (c cardMode) GetEnumValues() []string {
	return cardModes
}

type ownerMode struct{ *jsonschema.BasicEnum }

func (o ownerMode) GetEnumValues() []string {
	return ownerModes
}

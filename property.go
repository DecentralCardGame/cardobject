package cardobject

import "github.com/DecentralCardGame/jsonschema"

var actionIntProperties []string = []string{"COSTSUM"}
var actionStringProperties []string = []string{"NAME", "TEXT"}
var entityIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH"}
var entityStringProperties []string = []string{"NAME", "TEXT"}
var placeIntProperties []string = []string{"COSTSUM", "HEALTH"}
var placeStringProperties []string = []string{"NAME", "TEXT"}

var cardIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH"}
var cardStringProperties []string = []string{"NAME", "TEXT"}

var playerIntProperties []string = []string{"HANDSIZE", "BOARDSIZE", "DECKSIZE"}

type actionIntProperty struct{ *jsonschema.BasicEnum }

func (a actionIntProperty) GetEnumValues() []string {
	return actionIntProperties
}

type actionStringProperty struct{ *jsonschema.BasicEnum }

func (a actionStringProperty) GetEnumValues() []string {
	return actionStringProperties
}

type entityIntProperty struct{ *jsonschema.BasicEnum }

func (e entityIntProperty) GetEnumValues() []string {
	return entityIntProperties
}

type entityStringProperty struct{ *jsonschema.BasicEnum }

func (e entityStringProperty) GetEnumValues() []string {
	return entityStringProperties
}

type placeIntProperty struct{ *jsonschema.BasicEnum }

func (p placeIntProperty) GetEnumValues() []string {
	return placeIntProperties
}

type placeStringProperty struct{ *jsonschema.BasicEnum }

func (p placeStringProperty) GetEnumValues() []string {
	return actionStringProperties
}

type cardIntProperty struct{ *jsonschema.BasicEnum }

func (c cardIntProperty) GetEnumValues() []string {
	return cardIntProperties
}

type cardStringProperty struct{ *jsonschema.BasicEnum }

func (c cardStringProperty) GetEnumValues() []string {
	return cardStringProperties
}

type playerIntProperty struct{ *jsonschema.BasicEnum }

func (p playerIntProperty) GetEnumValues() []string {
	return playerIntProperties
}

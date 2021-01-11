package cardobject

import "github.com/DecentralCardGame/jsonschema"

var dynamicZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND"}
var zones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

var actionZones []string = []string{"DECK", "DUSTPILE", "HAND", "VOID"}
var entityZones []string = []string{"ATTACKLANE", "BLOCKLANE", "DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}
var placeZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

type dynamicZone struct{ *jsonschema.BasicEnum }

func (d dynamicZone) GetEnumValues() []string {
	return dynamicZones
}

type zone struct{ *jsonschema.BasicEnum }

func (z zone) GetEnumValues() []string {
	return zones
}

type actionZone struct{ *jsonschema.BasicEnum }

func (a actionZone) GetEnumValues() []string {
	return actionZones
}

type entityZone struct{ *jsonschema.BasicEnum }

func (e entityZone) GetEnumValues() []string {
	return entityZones
}

type placeZone struct{ *jsonschema.BasicEnum }

func (p placeZone) GetEnumValues() []string {
	return placeZones
}

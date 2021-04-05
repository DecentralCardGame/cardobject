package cardobject

import "github.com/DecentralCardGame/jsonschema"

type ActionZoneChange struct {
	Zone   ActionZone
	Player PlayerMode
}

func (a ActionZoneChange) Validate() error {
	return a.ValidateStruct()
}

func (a ActionZoneChange) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ActionZoneChange) InteractionText() string {
	return "Put it into §Player §Zone."
}

type EntityZoneChange struct {
	Zone   EntityZone
	Player PlayerMode
}

func (e EntityZoneChange) Validate() error {
	return e.ValidateStruct()
}

func (e EntityZoneChange) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e EntityZoneChange) InteractionText() string {
	return "Put it into §Player §Zone."
}

type PlaceZoneChange struct {
	Zone   PlaceZone
	Player PlayerMode
}

func (p PlaceZoneChange) Validate() error {
	return p.ValidateStruct()
}

func (p PlaceZoneChange) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p PlaceZoneChange) InteractionText() string {
	return "Put it into §Player §Zone."
}

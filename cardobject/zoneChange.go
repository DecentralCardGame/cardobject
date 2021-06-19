package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type ActionZoneChange struct {
	Zone    ActionZone
	Player  PlayerMode
	Keyword *Keyword `json:",omitempty"`
}

func (a ActionZoneChange) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ActionZoneChange) InteractionText() string {
	return "Put it into §Player §Zone.(§Keyword)"
}

type EntityZoneChange struct {
	Zone    EntityZone
	Player  PlayerMode
	Keyword *Keyword `json:",omitempty"`
}

func (e EntityZoneChange) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e EntityZoneChange) InteractionText() string {
	return "Put it into §Player §Zone.(§Keyword)"
}

type PlaceZoneChange struct {
	Zone    PlaceZone
	Player  PlayerMode
	Keyword *Keyword `json:",omitempty"`
}

func (p PlaceZoneChange) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p PlaceZoneChange) InteractionText() string {
	return "Put it into §Player §Zone.(§Keyword)"
}

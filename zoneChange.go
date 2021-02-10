package cardobject

import "github.com/DecentralCardGame/jsonschema"

type actionZoneChange struct {
	Zone   actionZone
	Player playerMode
}

func (a actionZoneChange) Validate() error {
	return a.ValidateStruct()
}

func (a actionZoneChange) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a actionZoneChange) GetInteractionText() string {
	return "Put it into §Player §Zone."
}

type entityZoneChange struct {
	Zone   entityZone
	Player playerMode
}

func (e entityZoneChange) Validate() error {
	return e.ValidateStruct()
}

func (e entityZoneChange) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entityZoneChange) GetInteractionText() string {
	return "Put it into §Player §Zone."
}

type placeZoneChange struct {
	Zone   placeZone
	Player playerMode
}

func (p placeZoneChange) Validate() error {
	return p.ValidateStruct()
}

func (p placeZoneChange) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p placeZoneChange) GetInteractionText() string {
	return "Put it into §Player §Zone."
}

package cardobject

import "github.com/DecentralCardGame/jsonschema"

type actionZoneChange struct {
	*jsonschema.BasicStruct
	Zone   actionZone
	Player playerMode
}

func (a actionZoneChange) GetInteractionText() string {
	return ""
}

type entityZoneChange struct {
	*jsonschema.BasicStruct
	Zone   entityZone
	Player playerMode
}

func (e entityZoneChange) GetInteractionText() string {
	return ""
}

type placeZoneChange struct {
	*jsonschema.BasicStruct
	Zone   placeZone
	Player playerMode
}

func (p placeZoneChange) GetInteractionText() string {
	return ""
}

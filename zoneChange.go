package cardobject

type actionZoneChange struct {
	Zone   string
	Player string
}

type entityZoneChange struct {
	Zone   string
	Player string
}

type placeZoneChange struct {
	Zone   string
	Player string
}

func (c *actionZoneChange) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionZone(c.Zone))
	errorRange = append(errorRange, validateOwnerMode(c.Player))
	return combineErrors(errorRange)
}

func (c *entityZoneChange) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityZone(c.Zone))
	errorRange = append(errorRange, validateOwnerMode(c.Player))
	return combineErrors(errorRange)
}

func (c *placeZoneChange) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceZone(c.Zone))
	errorRange = append(errorRange, validateOwnerMode(c.Player))
	return combineErrors(errorRange)
}

package cardobject

type Cost interface {
	GetRessources() []Ressource
	GetZoneChanges() []ZoneChange
}

func NewCost(r []Ressource, z []ZoneChange) Cost {
	return &cost{r, z}
}


type cost struct{
	RessourceCost []Ressource
	ZoneChangeCost []ZoneChange
}

func (c *cost) GetRessources() []Ressource {
	return c.RessourceCost
}

func (c *cost) GetZoneChanges() []ZoneChange {
	return c.ZoneChangeCost
}
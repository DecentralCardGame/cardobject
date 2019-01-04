package cardobject

type Cost interface {
	GetRessources() []Ressource
	GetZoneChanges() []ZoneChange
}

func NewCost(r []Ressource, z []ZoneChange) Cost {
	return &cost{r, z}
}


type cost struct{
	ressourceCost []Ressource
	zoneChangeCost []ZoneChange
}

func (c *cost) GetRessources() []Ressource {
	return c.ressourceCost
}

func (c *cost) GetZoneChanges() []ZoneChange {
	return c.zoneChangeCost
}
package cardobject

func NewCost(r []Ressource, z []ZoneChange) *Cost {
	return &Cost{r, z}
}


type Cost struct{
	ressourceCost []Ressource
	zoneChangeCost []ZoneChange
}

func (c *Cost) GetRessources() []Ressource {
	return c.ressourceCost
}

func (c *Cost) GetZoneChanges() []ZoneChange {
	return c.zoneChangeCost
}
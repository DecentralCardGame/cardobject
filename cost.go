package cardobject

type Cost interface {
	GetRessources() []Ressource
}

func NewCost(r []Ressource) Cost {
	return &cost{r}
}


type cost struct{
	RessourceCost []Ressource
}

func (c *cost) GetRessources() []Ressource {
	return c.RessourceCost
}
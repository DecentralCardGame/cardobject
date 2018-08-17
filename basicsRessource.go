package cardobject

type Ressource struct {
	rtype string 
}

func (r *Ressource) GetType() string{
	return r.rtype
}
 
func NewFoodRessource() *Ressource{
	return &Ressource{"FOOD"}
}

func NewGenericRessource() *Ressource{
	return &Ressource{"GENERIC"}
}

func NewLumberRessource() *Ressource{
	return &Ressource{"LUMBER"}
}

func NewManaRessource() *Ressource{
	return &Ressource{"MANA"}
}

func NewMetalRessource() *Ressource{
	return &Ressource{"METAL"}
}

func NewPowerRessource() *Ressource{
	return &Ressource{"POWER"}
}
package cardobject

const (
	ENERGY ressources = iota
	FOOD
	GENERIC
	LUMBER
	MANA
	METAL
)

type Ressource interface {
	Ressources() ressources
}


type ressources int

func (r ressources) Ressources() ressources {
	return r
}

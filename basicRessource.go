package cardobject

type ressources int

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

func (r ressources) Ressources() ressources {
	return r
}

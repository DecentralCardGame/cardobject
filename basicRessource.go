package cardobject

type ressources int

const (
	FOOD ressources = iota
	GENERIC
	LUMBER
	MANA
	METAL
	POWER
)

type Ressource interface {
	Ressources() ressources
}

func (r ressources) Ressources() ressources {
	return r
}

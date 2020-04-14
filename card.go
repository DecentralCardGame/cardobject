package cardobject

type cardAttributes struct {
	Name string
	Tags []string
	Text string
	CostType ressources
}

type action struct {
	cardAttributes
	CastingCost int8
	Effects []effect
}

type permanent struct {
	Abilities []ability
	Health int8
}

type entity struct {
	cardAttributes
	CastingCost int8
	permanent
	Attack int8
}

type place struct {
	cardAttributes
	CastingCost int8
	permanent
}

type headquarter struct {
	cardAttributes
	permanent
}

type card struct {
	Action *action `json:",omitempty"`
	Entity *entity `json:",omitempty"`
	Place *place `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

package cardobject

type cardAttributes struct {
	Name string
	Tags []string
	Text string
}

type castable struct {
	Cost ressources
	CastSpeed int8
}

type action struct {
	cardAttributes
	castable
	Effects []effect
}

type permanent struct {
	Abilities []ability
	AbilitySpeed int8
	Health int8
}

type entity struct {
	cardAttributes
	castable
	permanent
	Attack int8
}

type location struct {
	cardAttributes
	castable
	permanent
}

type headquarter struct {
	UniqueName string
	cardAttributes
	permanent
}

type card struct {
	Action *action `json:",omitempty"`
	Entity *entity `json:",omitempty"`
	Location *location `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

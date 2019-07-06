package cardobject

type card interface {
	getCardName() string
}

type cardAttributes struct {
	Name string
	Tag string
	Text string
}

type castable struct {
	Cost []string
	CastSpeed int8
}

type action struct {
	cardAttributes
	castable
	Effect effect
}

type permanent struct {
	Abilities []abilityWrapper
	AbilitySpeed int8
	Health int8
}

type entity struct {
	cardAttributes
	castable
	permanent
	Attack int8
}

type field struct {
	cardAttributes
	castable
	permanent	
}

type headquarter struct {
	cardAttributes
	permanent
	UniqueName string
}

func (ca cardAttributes) getCardName() string {
	return ca.Name
}
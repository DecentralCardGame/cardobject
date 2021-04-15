package keywords

import (
	"reflect"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

/*KeywordedCard is a keyworded card which can be resolved to a cardobject.Card
 */
type KeywordedCard interface {
	Resolve() cardobject.Card
}

type card struct {
	Action      *action      `json:",omitempty"`
	Entity      *entity      `json:",omitempty"`
	Place       *place       `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

func (c card) Resolve() cardobject.Card {
	valueOfB := reflect.ValueOf(c)
	typeOfB := reflect.TypeOf(c)
	possibleImplementer := []KeywordedCard{}
	for k := 0; k < valueOfB.NumField(); k++ {
		if !typeOfB.Field(k).Anonymous {
			possibleImplementer = append(possibleImplementer, valueOfB.Field(k).Interface().(KeywordedCard))
		}
	}
	for _, b := range possibleImplementer {
		if !reflect.ValueOf(b).IsNil() {
			return b.Resolve()
		}
	}
	return cardobject.Card{}
}

func (c card) Validate() error {
	return c.ValidateInterface()
}

func (c card) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type action struct {
	CardName    cardobject.CardName
	CastingCost cardobject.CastingCost
	CostType    cardobject.RessourceCostType
	Effects     effects
	FlavourText cardobject.FlavourText
	Tags        cardobject.Tags
}

func (a action) Resolve() cardobject.Card {
	effects := a.Effects.Resolve()
	card := cardobject.Card{
		Action: &cardobject.Action{
			CardName:    a.CardName,
			CastingCost: a.CastingCost,
			CostType:    a.CostType,
			Effects:     effects,
			FlavourText: a.FlavourText,
			Tags:        a.Tags}}
	return card
}

func (a action) Validate() error {
	return a.ValidateStruct()
}

func (a action) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a action) InteractionText() string {
	return "§CardName §CastingCost §CostType §Effects §FlavourText §Tags"
}

type entity struct {
	CardName    cardobject.CardName
	CastingCost cardobject.CastingCost
	CostType    cardobject.RessourceCostType
	Abilities   abilities
	Attack      cardobject.Attack
	Health      cardobject.Health
	FlavourText cardobject.FlavourText
	Tags        cardobject.Tags
}

func (e entity) Resolve() cardobject.Card {
	abilities := e.Abilities.Resolve()
	card := cardobject.Card{
		Entity: &cardobject.Entity{
			CardName:    e.CardName,
			CastingCost: e.CastingCost,
			CostType:    e.CostType,
			Abilities:   abilities,
			Attack:      e.Attack,
			Health:      e.Health,
			FlavourText: e.FlavourText,
			Tags:        e.Tags}}
	return card
}

func (e entity) Validate() error {
	return e.ValidateStruct()
}

func (e entity) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entity) InteractionText() string {
	return "§CardName §CastingCost §CostType §Abilities §Attack §Health §FlavourText §Tags"
}

type place struct {
	CardName    cardobject.CardName
	CastingCost cardobject.CastingCost
	CostType    cardobject.RessourceCostType
	Abilities   abilities
	Health      cardobject.Health
	FlavourText cardobject.FlavourText
	Tags        cardobject.Tags
}

func (p place) Resolve() cardobject.Card {
	abilities := p.Abilities.Resolve()
	card := cardobject.Card{
		Place: &cardobject.Place{
			CardName:    p.CardName,
			CastingCost: p.CastingCost,
			CostType:    p.CostType,
			Abilities:   abilities,
			Health:      p.Health,
			FlavourText: p.FlavourText,
			Tags:        p.Tags}}
	return card
}

func (p place) Validate() error {
	return p.ValidateStruct()
}

func (p place) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p place) InteractionText() string {
	return "§CardName §CastingCost §CostType §Abilities §Health §FlavourText §Tags"
}

type headquarter struct {
	CardName    cardobject.CardName
	CostType    cardobject.RessourceCostType
	Abilities   abilities
	Health      cardobject.Health
	FlavourText cardobject.FlavourText
	Tags        cardobject.Tags
}

func (h headquarter) Resolve() cardobject.Card {
	abilities := h.Abilities.Resolve()
	card := cardobject.Card{
		Headquarter: &cardobject.Headquarter{
			CardName:    h.CardName,
			CostType:    h.CostType,
			Abilities:   abilities,
			Health:      h.Health,
			FlavourText: h.FlavourText,
			Tags:        h.Tags}}
	return card
}

func (h headquarter) Validate() error {
	return h.ValidateStruct()
}

func (h headquarter) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h headquarter) InteractionText() string {
	return "§CardName §CostType §Abilities §Health §Growth §StartingHandSize §Wisdom §FlavourText §Tags"
}

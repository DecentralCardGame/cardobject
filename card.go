package cardobject

import "fmt"
import "encoding/json"

type Card interface {
	GetCost() Cost
	GetSpeed() Speedmodifier
	GetTags() []Tag
	Serialize() string
}

type Permanent interface {
	Card
	GetAbility() Ability
	GetHealth() Health
}

func NewAction(c Cost, s Speedmodifier, ts []Tag, t Text, e Effect) *Action {
	return &Action{&card{c, s, ts, t}, e}
}

func NewEntity(c Cost, s Speedmodifier, ts []Tag, t Text, a Ability, h Health, at Attack) *Entity {
	return &Entity{&permanent{&card{c, s, ts, t}, a, h}, at}
}

func NewField(c Cost, s Speedmodifier, ts []Tag, t Text, a Ability, h Health, st []Ressource) *Field {
	return &Field{&permanent{&card{c, s, ts, t}, a, h}, st, nil}
}


type card struct{
	cost Cost
	speedmodifier Speedmodifier
	tag []Tag
	text Text
}

type Action struct{
	*card
	effect Effect
}

type permanent struct {
	*card
	ability Ability
	health Health
}

type Entity struct{
	*permanent
	attack Attack
}

type Field struct{
	*permanent
	storage []Ressource
	inhabitants []Entity
}


func (c *card) GetCost() Cost {
	return c.cost
}

func (c *card) GetSpeed() Speedmodifier {
	return c.speedmodifier
}

func (c *card) GetTags() []Tag {
	return c.tag
}

func (c *card) Serialize() string {
	bytes, err := json.Marshal(c)
    if err != nil {
        fmt.Println("Can't serialize", c)
        return "Can't serialize"
    }
    return string(bytes)
}

func (a *Action) GetEffect() Effect {
	return a.effect
}

func (p *permanent) GetAbility() Ability {
	return p.ability
}

func (p *permanent) GetHealth() Health {
	return p.health
}

func (e Entity) GetAttack() Attack {
	return e.attack
}

func (f *Field) GetStorage() []Ressource {
	return f.storage
}
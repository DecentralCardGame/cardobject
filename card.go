package cardobject

import "fmt"
import "encoding/json"

type Card interface {
	GetCost() Cost
	GetSpeedModifier() SpeedModifier
	GetTags() []Tag
	GetText() Text
	Serialize() string
}

type Action interface {
	Card
	GetEffect() Effect
}

type Permanent interface {
	Card
	GetAbility() Ability
	GetHealth() Health
}

type Entity interface {
	Permanent
	GetAttack() Attack
}

type Field interface {
	Permanent
}

func NewAction(c Cost, s SpeedModifier, ts []Tag, t Text, e Effect) Action {
	return &action{&card{c, s, ts, t}, e}
}

func NewEntity(c Cost, s SpeedModifier, ts []Tag, t Text, a Ability, h Health, at Attack) Entity {
	return &entity{&permanent{&card{c, s, ts, t}, a, h}, at}
}

func NewField(c Cost, s SpeedModifier, ts []Tag, t Text, a Ability, h Health) Field {
	return &field{&permanent{&card{c, s, ts, t}, a, h}}
}


type card struct{
	cost Cost
	speedmodifier SpeedModifier
	tag []Tag
	text Text
}

type action struct{
	*card
	effect Effect
}

type permanent struct {
	*card
	ability Ability
	health Health
}

type entity struct{
	*permanent
	attack Attack
}

type field struct{
	*permanent
}


func (c *card) GetCost() Cost {
	return c.cost
}

func (c *card) GetSpeedModifier() SpeedModifier {
	return c.speedmodifier
}

func (c *card) GetTags() []Tag {
	return c.tag
}

func (c *card) GetText() Text {
	return c.text
}

func (c *card) Serialize() string {
	bytes, err := json.Marshal(c)
    if err != nil {
        fmt.Println("Can't serialize", c)
        return "Can't serialize"
    }
    return string(bytes)
}

func (a *action) GetEffect() Effect {
	return a.effect
}

func (p *permanent) GetAbility() Ability {
	return p.ability
}

func (p *permanent) GetHealth() Health {
	return p.health
}

func (e entity) GetAttack() Attack {
	return e.attack
}
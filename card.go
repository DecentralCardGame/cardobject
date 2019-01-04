package cardobject

type Card interface {
	GetCost() Cost
	GetSpeedModifier() SpeedModifier
	GetTags() []Tag
	GetText() Text
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
	Cost Cost
	Speedmodifier SpeedModifier
	Tag []Tag
	Text Text
}

type action struct{
	*card
	Effect Effect
}

type permanent struct {
	*card
	Ability Ability
	Health Health
}

type entity struct{
	*permanent
	Attack Attack
}

type field struct{
	*permanent
}


func (c *card) GetCost() Cost {
	return c.Cost
}

func (c *card) GetSpeedModifier() SpeedModifier {
	return c.Speedmodifier
}

func (c *card) GetTags() []Tag {
	return c.Tag
}

func (c *card) GetText() Text {
	return c.Text
}

func (a *action) GetEffect() Effect {
	return a.Effect
}

func (p *permanent) GetAbility() Ability {
	return p.Ability
}

func (p *permanent) GetHealth() Health {
	return p.Health
}

func (e entity) GetAttack() Attack {
	return e.Attack
}
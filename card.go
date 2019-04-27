package cardobject

type Card interface {
	GetName() Name
	GetTags() []Tag
	GetText() Text
}

type Castable interface {
	GetCost() Cost
	GetCastSpeed() SpeedModifier
}

type Permanent interface {
	GetAbility() Ability
	GetAbilitySpeed() SpeedModifier
	GetHealth() Health
}

type Action interface {
	Card
	Castable
	GetEffect() Effect
}

type Entity interface {
	Card
	Castable
	Permanent
	GetAttack() Attack
}

type Field interface {
	Card
	Castable
	Permanent
}

type Headquarter interface {
	Card
	Permanent
	GetUniqueName() UniqueName
}

func NewAction(n Name, c Cost, s SpeedModifier, ts []Tag, t Text, e Effect) Action {
	return &action{&cardAttributes{n, ts, t}, &castable{c, s}, e}
}

func NewEntity(n Name, c Cost, cs SpeedModifier, ts []Tag, t Text, a Ability, as SpeedModifier, h Health, at Attack) Entity {
	return &entity{&cardAttributes{n, ts, t}, &castable{c, cs}, &permanent{a, as, h}, at}
}

func NewField(n Name, c Cost, cs SpeedModifier, ts []Tag, t Text, a Ability, as SpeedModifier, h Health) Field {
	return &field{&cardAttributes{n, ts, t}, &castable{c, cs}, &permanent{a, as, h}}
}

func NewHeadquarter(un UniqueName, ts []Tag, t Text, a Ability, as SpeedModifier, h Health) Headquarter {
	return &headquarter{&cardAttributes{NewName(un.GetStringVal()), ts, t}, &permanent{a, as, h}, un}
}

type cardAttributes struct {
	Name Name
	Tag []Tag
	Text Text
}

type castable struct {
	Cost Cost
	CastSpeed SpeedModifier
}

type action struct {
	*cardAttributes
	*castable
	Effect Effect
}

type permanent struct {
	Ability Ability
	AbilitySpeed SpeedModifier
	Health Health
}

type entity struct {
	*cardAttributes
	*castable
	*permanent
	Attack Attack
}

type field struct {
	*cardAttributes
	*castable
	*permanent	
}

type headquarter struct {
	*cardAttributes
	*permanent
	UniqueName UniqueName
}

func (c *cardAttributes) GetName() Name {
	return c.Name
}

func (c *castable) GetCost() Cost {
	return c.Cost
}

func (c *castable) GetCastSpeed() SpeedModifier {
	return c.CastSpeed
}

func (c *cardAttributes) GetTags() []Tag {
	return c.Tag
}

func (c *cardAttributes) GetText() Text {
	return c.Text
}

func (a *action) GetEffect() Effect {
	return a.Effect
}

func (p *permanent) GetAbility() Ability {
	return p.Ability
}

func (p *permanent) GetAbilitySpeed() SpeedModifier {
	return p.AbilitySpeed
}

func (p *permanent) GetHealth() Health {
	return p.Health
}

func (e entity) GetAttack() Attack {
	return e.Attack
}

func (hq headquarter) GetUniqueName() UniqueName {
	return hq.UniqueName
}
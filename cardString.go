package cardobject

type readableAction struct {
	cardAttributes
	castable
	Effect effect
	ReadAbleEffect string
}

type readablePermanent struct {
	Abilities []abilityWrapper
	ReadableAbilities string
	AbilitySpeed int8
	Health int8
}

type readableEntity struct {
	cardAttributes
	castable
	readablePermanent
	Attack int8
}

type readableField struct {
	cardAttributes
	castable
	readablePermanent	
}

type readableHeadquarter struct {
	cardAttributes
	readablePermanent
	UniqueName string
}

type readableCardWrapper struct {
	Action *readableAction `json:",omitempty"`
	Entity *readableEntity `json:",omitempty"`
	Field *readableField `json:",omitempty"`
	Headquarter *readableHeadquarter `json:",omitempty"`
}

func (cw *cardWrapper) toReadable() readableCardWrapper {
	a := cw.Action
	e := cw.Entity
	f := cw.Field
	h := cw.Headquarter
	if(a != nil) {
		return readableCardWrapper{a.toReadable(), nil, nil, nil}
	}
	if(e != nil) {
		return readableCardWrapper{nil, e.toReadable(), nil, nil}
	}
	if(f != nil) {
		return readableCardWrapper{nil, nil, f.toReadable(), nil}
	}
	if(h != nil) {
		return readableCardWrapper{nil, nil, nil, h.toReadable()}
	}
	return readableCardWrapper{}
}

func (a *action) toReadable() *readableAction {
	return &readableAction{cardAttributes{a.Name, a.Tag, a.Text}, castable{a.Cost, a.CastSpeed}, a.Effect, a.Effect.ToString()}
}

func (e *entity) toReadable() *readableEntity {
	return &readableEntity{cardAttributes{e.Name, e.Tag, e.Text}, castable{e.Cost, e.CastSpeed}, readablePermanent{e.Abilities, ToString(e.Abilities), e.AbilitySpeed, e.Health}, e.Attack}
}

func (f *field) toReadable() *readableField {
	return &readableField{cardAttributes{f.Name, f.Tag, f.Text}, castable{f.Cost, f.CastSpeed}, readablePermanent{f.Abilities, ToString(f.Abilities), f.AbilitySpeed, f.Health}}
}

func (h *headquarter) toReadable() *readableHeadquarter {
	return &readableHeadquarter{cardAttributes{h.Name, h.Tag, h.Text}, readablePermanent{h.Abilities, ToString(h.Abilities), h.AbilitySpeed, h.Health}, h.UniqueName}
}
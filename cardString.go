package cardobject

import "strings"

type readable struct {
	ReadAbleCardText string
}

type readableCard struct {
	card
	readable
}

func (c *card) toReadable() readableCard {
	a := c.Action
	e := c.Entity
	f := c.Place
	h := c.Headquarter
	if(a != nil) {
		return readableCard{*c, readable{a.toString()}}
	}
	if(e != nil) {
		return readableCard{*c, readable{e.toString()}}
	}
	if(f != nil) {
		return readableCard{*c, readable{f.toString()}}
	}
	if(h != nil) {
		return readableCard{*c, readable{h.toString()}}
	}
	return readableCard{}
}

func (a *action) toString() string {
	var effectsString []string
	effects := a.Effects
	if(effects != nil) {
		for _, e := range effects {
    		effectsString = append(effectsString, e.toString())
		}
	}
	return strings.Join(effectsString, "\n")
}

func (e *entity) toString() string {
	var abilitiesString []string
	abilities := e.Abilities
	if(abilities != nil) {
		for _, a := range abilities {
    		abilitiesString = append(abilitiesString, a.toString())
		}
	}
	return strings.Join(abilitiesString, "\n")
}

func (f *place) toString() string {
	var abilitiesString []string
	abilities := f.Abilities
	if(abilities != nil) {
		for _, a := range abilities {
    		abilitiesString = append(abilitiesString, a.toString())
		}
	}
	return strings.Join(abilitiesString, "\n")
}

func (h *headquarter) toString() string {
	var abilitiesString []string
	abilities := h.Abilities
	if(abilities != nil) {
		for _, a := range abilities {
    		abilitiesString = append(abilitiesString, a.toString())
		}
	}
	return strings.Join(abilitiesString, "\n")
}

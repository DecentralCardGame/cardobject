package cardobject

import "strings"

func (a *ability) toString() string {
	var plainText string
	aa := a.ActivatedAbility
	ta := a.TriggeredAbility
	if(aa != nil) {
		return aa.toString()
	}
	if(ta != nil) {
		return ta.toString()
	}
	return plainText
}

func (aa *activatedAbility) toString() string {
	var plainText string
	effects := aa.Effects

	if(aa.MultipleUse) {
		plainText += "|M|"
	}

	plainText += "Pay " + aa.Cost.toString() + ": "

	var effectsString []string
	if(effects != nil) {
		for _, e := range effects {
    		effectsString = append(effectsString, e.toString())
		}
	}
	plainText += strings.Join(effectsString, " ")

	return plainText
}

func (ta *triggeredAbility) toString() string {
	var plainText string
	effects := ta.Effects
	plainText += ta.Cause.toString()

	var effectsString []string
	if(effects != nil) {
		for _, e := range effects {
    		effectsString = append(effectsString, e.toString())
		}
	}
	plainText += strings.Join(effectsString, " ")

	return plainText
}
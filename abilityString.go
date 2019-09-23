package cardobject

import "strings"

func ToString(aws []abilityWrapper) string {
	var plainText string
	for i, a := range aws {
			if(i > 0) {
				plainText += "\n"
			}
    		plainText += a.ToString()
	}
	return plainText
}

func (aw *abilityWrapper) ToString() string {
	var plainText string
	aa := aw.ActivatedAbility
	ta := aw.TriggeredAbility
	if(aa != nil) {
		return aa.ToString()
	}
	if(ta != nil) {
		return ta.ToString()
	}
	return plainText
}

func (aa *activatedAbility) ToString() string {
	var plainText string
	if(aa.MultipleUse) {
		plainText += "|M|"
	}
	plainText += "Pay " + strings.Join(aa.Cost, ",") + ": "
	plainText += aa.Effect.ToString()
	return plainText
}

func (ta *triggeredAbility) ToString() string {
	var plainText string
	plainText += ta.Cause.ToString()
	plainText += ta.Effect.ToString()
	return plainText
}
package cardobject

import "strings"

func (aw *abilityWrapper) ToPlainText() string {
	var plainText string
	aa := aw.ActivatedAbility
	ta := aw.TriggeredAbility
	if(aa != nil) {
		return aa.ToPlainText()
	}
	if(ta != nil) {
		return ta.ToPlainText()
	}
	return plainText
}

func (aa *activatedAbility) ToPlainText() string {
	var plainText string
	if(aa.MultipleUse) {
		plainText += "|M|"
	}
	plainText += "Pay " + strings.Join(aa.Cost, ",") + ": "
	plainText += aa.Effect.ToPlainText()
	return plainText
}

func (ta *triggeredAbility) ToPlainText() string {
	var plainText string
	plainText += ta.Cause.ToPlainText()
	plainText += ta.Effect.ToPlainText()
	return plainText
}
package cardobject

import "strings"
import "strconv"

func (e *effect) toString() string {
	var sentences []string
	produce := e.Production
	draw := e.Draw
	tokenEffect := e.TokenEffect
	targetEffect := e.TargetEffect

	if (targetEffect != nil) {
		sentences = append(sentences, targetEffect.toString())
	}

	if (tokenEffect != nil) {
		sentences = append(sentences, tokenEffect.toString())
	}

	if(len(produce) > 0) {
		sentences = append(sentences, "Produce " + strings.Join(produce, ",") + ".")
	}

	if(draw != nil) {
		sentences = append(sentences, "Draw " + strconv.Itoa(*draw) + " cards.")
	}

	return strings.Join(sentences, " ")
}

func (te *targetEffect) toString() string {
	ate := te.ActionTargetEffect
	ete := te.EntityTargetEffect
	fte := te.FieldTargetEffect
	if(ate != nil) {
		return ate.toString()
	}
	if(ete != nil) {
		return ete.toString()
	}
	if(fte != nil) {
		return fte.toString()
	}
	return ""
}

func (ate *actionTargetEffect) toString() string {
	var sentences []string
	plural := false
	manipulations := ate.ActionManipulations
	zoneChange := ate.ZoneChange
	if(ate.ActionSelector.CardMode == "ALL") {
		plural = true
	}

	sentences = append(sentences, ate.ActionSelector.toString())
	
	if(manipulations != nil) {
		for _, m := range manipulations {
    		sentences = append(sentences, m.toString(plural))
		}
	}

	if(zoneChange != nil) {
		sentences = append(sentences, zoneChangeString(zoneChange, plural))
	}
	return strings.Join(sentences, " ")
}

func (ete *entityTargetEffect) toString() string {
	var sentences []string
	plural := false
	manipulations := ete.EntityManipulations
	zoneChange := ete.ZoneChange
	if(ete.EntitySelector.CardMode == "ALL") {
		plural = true
	}

	sentences = append(sentences, ete.EntitySelector.toString())
	
	if(manipulations != nil) {
		for _, m := range manipulations {
    		sentences = append(sentences, m.toString(plural))
		}
	}

	if(zoneChange != nil) {
		sentences = append(sentences, zoneChangeString(zoneChange, plural))
	}
	return strings.Join(sentences, " ")
}

func (fte *fieldTargetEffect) toString() string {
	var sentences []string
	plural := false
	manipulations := fte.FieldManipulations
	zoneChange := fte.ZoneChange
	if(fte.FieldSelector.CardMode == "ALL") {
		plural = true
	}

	sentences = append(sentences, fte.FieldSelector.toString())
	
	if(manipulations != nil) {
		for _, m := range manipulations {
    		sentences = append(sentences, m.toString(plural))
		}
	}

	if(zoneChange != nil) {
		sentences = append(sentences, zoneChangeString(zoneChange, plural))
	}
	return strings.Join(sentences, " ")
}

func zoneChangeString(zone *string, plural bool) string {
	if(plural) {
		return "Move them to the " + *zone + "."
	}
	return "Move it to the " + *zone + "."
}

func (toe *tokenEffect) toString() string {
	var plainText string
	plainText += "Create "
	if (toe.Amount != nil) {
		plainText += strconv.Itoa(int(*toe.Amount))
	} else {
		plainText += "a"
	}
	plainText += " "
	plainText += toe.Token.toString()
	plainText += "."
	return plainText
}
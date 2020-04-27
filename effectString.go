package cardobject

import (
	"strconv"
	"strings"
)

func (e *effect) toString() string {
	var sentences []string
	productionEffect := e.ProductionEffect
	draw := e.Draw
	tokenEffect := e.TokenEffect
	targetEffect := e.TargetEffect

	if targetEffect != nil {
		sentences = append(sentences, targetEffect.toString())
	}

	if tokenEffect != nil {
		sentences = append(sentences, tokenEffect.toString())
	}

	if productionEffect != nil {
		sentences = append(sentences, tokenEffect.toString())
	}

	if draw != nil {
		sentences = append(sentences, "Draw "+strconv.Itoa(*draw)+" cards.")
	}

	return strings.Join(sentences, " ")
}

func (te *targetEffect) toString() string {
	ate := te.ActionTargetEffect
	ete := te.EntityTargetEffect
	fte := te.PlaceTargetEffect
	if ate != nil {
		return ate.toString()
	}
	if ete != nil {
		return ete.toString()
	}
	if fte != nil {
		return fte.toString()
	}
	return ""
}

func (pe *productionEffect) toString() string {
	return "Produce " + strconv.Itoa(int(pe.ProductionAmount)) + " of ."
}

func (ate *actionTargetEffect) toString() string {
	var sentences []string
	plural := false
	manipulations := ate.ActionManipulations
	zoneChange := ate.ZoneChange
	if ate.ActionSelector.CardMode == "ALL" {
		plural = true
	}

	sentences = append(sentences, ate.ActionSelector.toString())

	if manipulations != nil {
		for _, m := range manipulations {
			sentences = append(sentences, m.toString(plural))
		}
	}

	if zoneChange != nil {
		sentences = append(sentences, zoneChangeString(zoneChange, plural))
	}
	return strings.Join(sentences, " ")
}

func (ete *entityTargetEffect) toString() string {
	var sentences []string
	plural := false
	manipulations := ete.EntityManipulations
	zoneChange := ete.ZoneChange
	if ete.EntitySelector.CardMode == "ALL" {
		plural = true
	}

	sentences = append(sentences, ete.EntitySelector.toString())

	if manipulations != nil {
		for _, m := range manipulations {
			sentences = append(sentences, m.toString(plural))
		}
	}

	if zoneChange != nil {
		sentences = append(sentences, zoneChangeString(zoneChange, plural))
	}
	return strings.Join(sentences, " ")
}

func (fte *placeTargetEffect) toString() string {
	var sentences []string
	plural := false
	manipulations := fte.PlaceManipulations
	zoneChange := fte.ZoneChange
	if fte.PlaceSelector.CardMode == "ALL" {
		plural = true
	}

	sentences = append(sentences, fte.PlaceSelector.toString())

	if manipulations != nil {
		for _, m := range manipulations {
			sentences = append(sentences, m.toString(plural))
		}
	}

	if zoneChange != nil {
		sentences = append(sentences, zoneChangeString(zoneChange, plural))
	}
	return strings.Join(sentences, " ")
}

func zoneChangeString(zoneChange *zoneChange, plural bool) string {
	if plural {
		return "Move them to " + zoneChange.Player + " " + zoneChange.Zone + "."
	}
	return "Move it to the " + zoneChange.Player + " " + zoneChange.Zone + "."
}

func (toe *tokenEffect) toString() string {
	var plainText string
	plainText += "Create "
	if toe.Amount != nil {
		plainText += strconv.Itoa(int(*toe.Amount))
	} else {
		plainText += "a"
	}
	plainText += " "
	plainText += toe.Token.toString()
	plainText += "."
	return plainText
}

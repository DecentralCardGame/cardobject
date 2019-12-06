package cardobject

import "strconv"

func (m *actionManipulation) toString(plural bool) string {
	var plainText string
	im := m.ActionIntManipulation
	sm := m.ActionStringManipulation

	if(im != nil) {
		return im.toString(plural)
	}
	if(sm != nil) {
		return sm.toString(plural)
	}
	return plainText
}

func (m *entityManipulation) toString(plural bool) string {
	var plainText string
	am := m.EntityAbilityManipulation
	im := m.EntityIntManipulation
	sm := m.EntityStringManipulation

	if(am != nil) {
		return am.toString(plural)
	}
	if(im != nil) {
		return im.toString(plural)
	}
	if(sm != nil) {
		return sm.toString(plural)
	}
	return plainText
}

func (m *locationManipulation) toString(plural bool) string {
	var plainText string
	am := m.LocationAbilityManipulation
	im := m.LocationIntManipulation
	sm := m.LocationStringManipulation

	if(am != nil) {
		return am.toString(plural)
	}
	if(im != nil) {
		return im.toString(plural)
	}
	if(sm != nil) {
		return sm.toString(plural)
	}
	return plainText
}

func (am *abilityManipulationBasics) toString(plural bool) string {
	var plainText string
	if (plural) {
		plainText += "They "
	} else {
		plainText += "It "
	}
	plainText += am.AbilityOperator
	if (!plural) {
		plainText += "s "
	}
	plainText += am.Ability.toString()
	return plainText
}

func (im *intManipulationBasics) toString(plural bool) string {
	var plainText string

	switch im.IntOperator {
	case "SET":
		plainText += "Set "
		if (plural) {
			plainText += "their "
		} else {
			plainText += "its "
		}
		plainText += im.IntProperty + " to " + strconv.Itoa(im.IntValue) + "."
	case "ADD":
		plainText += "Add " + strconv.Itoa(im.IntValue) + " to "
		if (plural) {
			plainText += "their "
		} else {
			plainText += "its "
		}
		plainText +=  im.IntProperty + "."
	case "SUBTRACT":
		plainText += "Subtract " + strconv.Itoa(im.IntValue) + " from "
		if (plural) {
			plainText += "their "
		} else {
			plainText += "its "
		}
		plainText += im.IntProperty + "."
	default:

	}

	return plainText
}

func (sm *stringManipulationBasics) toString(plural bool) string {
	var plainText string

	plainText += "Set "
	if (plural) {
		plainText += "their "
	} else {
		plainText += "its "
	}
	plainText += sm.StringProperty + " to " + sm.StringValue + "."
	return plainText
}

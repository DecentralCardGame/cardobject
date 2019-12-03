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
	im := m.EntityIntManipulation
	sm := m.EntityStringManipulation

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
	im := m.LocationIntManipulation
	sm := m.LocationStringManipulation

	if(im != nil) {
		return im.toString(plural)
	}
	if(sm != nil) {
		return sm.toString(plural)
	}
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

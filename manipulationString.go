package cardobject

import "strconv"

func (mw *manipulationWrapper) ToPlainText() string {
	var plainText string
	im := mw.IntManipulation
	sm := mw.StringManipulation
	if(im != nil) {
		return im.ToPlainText()
	}
	if(sm != nil) {
		return im.ToPlainText()
	}
	return plainText
}

func (im *intManipulation) ToPlainText() string {
	var plainText string
	selectorText := im.Selector.ToPlainText()
	switch im.Operator {
	case "SET":
		plainText += "Set the " + im.Property + " of " + selectorText + " to " + strconv.Itoa(im.Value)
	case "ADD":
		plainText += "Add " + strconv.Itoa(im.Value) + " to the " + im.Property + " of " + selectorText
	case "SUBTRACT":
		plainText += "Subtract " + strconv.Itoa(im.Value) + " from the " + im.Property + " of " + selectorText
	default:
		
	}
	plainText += " " + im.Duration + "."
	return plainText
}

func (sm *stringManipulation) ToPlainText() string {
	var plainText string
	selectorText := sm.Selector.ToPlainText()
	plainText += "Set the " + sm.Property + " of " + selectorText + " to " + sm.Value + " " + sm.Duration
	return plainText
}

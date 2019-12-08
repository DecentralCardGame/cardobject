package cardobject

func (el *eventListener) toString() string {
	var plainText string
	pel := el.ProductionEventListener
	tel := el.TimeEventListener
	mel := el.ManipulationEventListener
	zel := el.ZoneChangeEventListener
	if(pel != nil) {
		return pel.toString()
	}
	if(tel != nil) {
		return tel.toString()
	}
	if(mel != nil) {
		return mel.toString()
	}
	if(zel != nil) {
		return zel.toString()
	}
	return plainText
}

func (pel *productionEventListener) toString() string {
	var plainText string
	plainText += "Whenever you produce " + pel.RessourceType + ": "
	return plainText
}

func (tel *timeEventListener) toString() string {
	var plainText string
	plainText += "At " + tel.TimeEvent + ": "
	return plainText
}

func (mel *manipulationEventListener) toString() string {
	var plainText string
	intMel := mel.IntManipulationEventListener
	stringMel := mel.StringManipulationEventListener
	plainText += "Whenever the "
	if(intMel != nil) {
		plainText += intMel.Property
	}
	if(stringMel != nil) {
		plainText += stringMel.Property
	}
	plainText += " of a card "
	if(intMel != nil) {
		plainText += intMel.ChangeMode
	}
	if(stringMel != nil) {
		plainText += stringMel.ChangeMode
	}
	plainText += ": "
	return plainText
}

func (zel *zoneChangeEventListener) toString() string {
	var plainText string
	plainText += "Whenever a card is moved from the "
	plainText += zel.Source+ " to the " + zel.Destination + ": "
	return plainText
}
package cardobject

func (el *eventListener) toString() string {
	var plainText string
	tel := el.TimeEventListener
	mel := el.ManipulationEventListener
	zel := el.ZoneChangeEventListener
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

func (tel *timeEventListener) toString() string {
	var plainText string
	plainText += "At " + tel.TimeEvent + ": "
	return plainText
}

func (mel *manipulationEventListener) toString() string {
	var plainText string
	plainText += "Whenever the " + mel.Property + " of a card changes: "
	return plainText
}

func (zel *zoneChangeEventListener) toString() string {
	var plainText string
	plainText += "Whenever a card is moved from the "
	plainText += zel.Source+ " to the " + zel.Destination + ": "
	return plainText
}
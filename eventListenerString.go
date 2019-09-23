package cardobject

func (elw *eventListenerWrapper) ToString() string {
	var plainText string
	tel := elw.TimeEventListener
	mel := elw.ManipulationEventListener
	zel := elw.ZoneChangeEventListener
	if(tel != nil) {
		return tel.ToString()
	}
	if(mel != nil) {
		return mel.ToString()
	}
	if(zel != nil) {
		return zel.ToString()
	}
	return plainText
}

func (tel *timeEventListener) ToString() string {
	var plainText string
	plainText += "At " + tel.TimeEvent + ": "
	return plainText
}

func (mel *manipulationEventListener) ToString() string {
	var plainText string
	plainText += "Whenever the " + mel.Property + " of a card changes: "
	return plainText
}

func (zel *zoneChangeEventListener) ToString() string {
	var plainText string
	plainText += "Whenever a card is moved from the "
	plainText += zel.Source+ " to the " + zel.Destination + ": "
	return plainText
}
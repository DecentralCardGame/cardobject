package cardobject

func (elw *eventListenerWrapper) ToPlainText() string {
	var plainText string
	tel := elw.TimeEventListener
	mel := elw.ManipulationEventListener
	zel := elw.ZoneChangeEventListener
	if(tel != nil) {
		return tel.ToPlainText()
	}
	if(mel != nil) {
		return mel.ToPlainText()
	}
	if(zel != nil) {
		return zel.ToPlainText()
	}
	return plainText
}

func (tel *timeEventListener) ToPlainText() string {
	var plainText string
	plainText += "At " + tel.TimeEvent + ": "
	return plainText
}

func (mel *manipulationEventListener) ToPlainText() string {
	var plainText string
	plainText += "Whenever the " + mel.Property + " of a card changes: "
	return plainText
}

func (zel *zoneChangeEventListener) ToPlainText() string {
	var plainText string
	plainText += "Whenever a card is moved from the "
	plainText += zel.Source+ " to the " + zel.Destination + ": "
	return plainText
}
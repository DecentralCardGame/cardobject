package cardobject

func (as *actionSelector) toString() string {
	plainText := "Choose " + as.CardMode + " action"
	if(as.CardMode == "ALL") {
		plainText += "s"
	}
	if(as.CardCondition != nil) {
		plainText += " with " + as.CardCondition.ToString()
	}
	plainText += " in the " + as.Zone
	if(as.PlayerMode == "ALL") {
		plainText += "s"
	}
	plainText += " of " + as.PlayerMode + " player"
	if(as.PlayerMode == "ALL") {
		plainText += "s"
	}
	if(as.PlayerCondition != nil) {
		plainText += " with " + as.PlayerCondition.ToString()
	}
	plainText += "."
	return plainText
}

func (es *entitySelector) toString() string {
	plainText := "Choose " + es.CardMode
	if(es.CardMode == "ALL") {
		plainText += " entities"
	} else {
		plainText += " entity"
	}
	if(es.CardCondition != nil) {
		plainText += " with " + es.CardCondition.ToString()
	}
	plainText += " in the " + es.Zone
	if(es.PlayerMode == "ALL") {
		plainText += "s"
	}
	plainText += " of " + es.PlayerMode + " player"
	if(es.PlayerMode == "ALL") {
		plainText += "s"
	}
	if(es.PlayerCondition != nil) {
		plainText += " with " + es.PlayerCondition.ToString()
	}
	plainText += "."
	return plainText
}

func (fs *placeSelector) toString() string {
	plainText := "Choose " + fs.CardMode + " place"
	if(fs.CardMode == "ALL") {
		plainText += "s"
	}
	if(fs.CardCondition != nil) {
		plainText += " with " + fs.CardCondition.ToString()
	}
	plainText += " in the " + fs.Zone
	if(fs.PlayerMode == "ALL") {
		plainText += "s"
	}
	plainText += " of " + fs.PlayerMode + " player"
	if(fs.PlayerMode == "ALL") {
		plainText += "s"
	}
	if(fs.PlayerCondition != nil) {
		plainText += " with " + fs.PlayerCondition.ToString()
	}
	plainText += "."
	return plainText
}

func (ss *selfSelector) toString() string {
	plainText := ss.Target
	return plainText
}

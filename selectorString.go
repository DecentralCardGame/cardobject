package cardobject

func (s *selectorWrapper) ToString() string {
	var plainText string
	as := s.ActionSelector
	es := s.EntitySelector
	fs := s.FieldSelector
	ss := s.SelfSelector
	if(as != nil) {
		return as.ToString()
	}
	if(es != nil) {
		return es.ToString()
	}
	if(fs != nil) {
		return fs.ToString()
	}
	if(ss != nil) {
		return ss.ToString()
	}
	return plainText
}

func (as *actionSelector) ToString() string {
	plainText := as.CardMode + " action"
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
	return plainText
}

func (es *entitySelector) ToString() string {
	plainText := es.CardMode
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
	return plainText
}

func (fs *fieldSelector) ToString() string {
	plainText := fs.CardMode + " field"
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
	return plainText
}

func (ss *selfSelector) ToString() string {
	plainText := ss.Target
	return plainText
}


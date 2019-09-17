package cardobject

func (s *selectorWrapper) ToPlainText() string {
	var plainText string
	as := s.ActionSelector
	es := s.EntitySelector
	fs := s.FieldSelector
	ss := s.SelfSelector
	if(as != nil) {
		return as.ToPlainText()
	}
	if(es != nil) {
		return es.ToPlainText()
	}
	if(fs != nil) {
		return fs.ToPlainText()
	}
	if(ss != nil) {
		return ss.ToPlainText()
	}
	return plainText
}

func (as *actionSelector) ToPlainText() string {
	plainText := as.CardMode + " action"
	if(as.CardMode == "ALL") {
		plainText += "s"
	}
	if(as.CardCondition != nil) {
		plainText += " with " + as.CardCondition.ToPlainText()
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
		plainText += " with " + as.PlayerCondition.ToPlainText()
	}
	return plainText
}

func (es *entitySelector) ToPlainText() string {
	plainText := es.CardMode
	if(es.CardMode == "ALL") {
		plainText += " entities"
	} else {
		plainText += " entity"
	}
	if(es.CardCondition != nil) {
		plainText += " with " + es.CardCondition.ToPlainText()
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
		plainText += " with " + es.PlayerCondition.ToPlainText()
	}
	return plainText
}

func (fs *fieldSelector) ToPlainText() string {
	plainText := fs.CardMode + " field"
	if(fs.CardMode == "ALL") {
		plainText += "s"
	}
	if(fs.CardCondition != nil) {
		plainText += " with " + fs.CardCondition.ToPlainText()
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
		plainText += " with " + fs.PlayerCondition.ToPlainText()
	}
	return plainText
}

func (ss *selfSelector) ToPlainText() string {
	plainText := ss.Target
	return plainText
}


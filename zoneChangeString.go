package cardobject

func (z *zoneChange) ToPlainText() string {
	var plainText string
	plainText += "Move " + z.Selector.ToPlainText() + " to the " + z.Destination + ". "
	return plainText
}
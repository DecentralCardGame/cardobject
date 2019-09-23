package cardobject

func (z *zoneChange) ToString() string {
	var plainText string
	plainText += "Move " + z.Selector.ToString() + " to the " + z.Destination + ". "
	return plainText
}
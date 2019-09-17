package cardobject

import "strconv"

func (cw *conditionWrapper) ToPlainText() string {
	var plainText string
	ic := cw.IntCondition
	sc := cw.StringCondition
	if(ic != nil) {
		return ic.ToPlainText()
	}
	if(sc != nil) {
		return sc.ToPlainText()
	}
	return plainText
}

func (ic *intCondition) ToPlainText() string {
	plainText := ic.Property + " " + ic.IntComparator + " " + strconv.Itoa(ic.IntValue)
	return plainText
}

func (sc *stringCondition) ToPlainText() string {
	plainText := sc.Property + " " + sc.StringComparator + " " + sc.StringValue
	return plainText
}




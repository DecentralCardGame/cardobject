package cardobject

import "strconv"

func (cw *conditionAttributes) ToString() string {
	var plainText string
	ic := cw.IntCondition
	sc := cw.StringCondition
	if(ic != nil) {
		return ic.ToString()
	}
	if(sc != nil) {
		return sc.ToString()
	}
	return plainText
}

func (ic *intCondition) ToString() string {
	plainText := ic.IntProperty + " " + ic.IntComparator + " " + strconv.Itoa(ic.IntValue)
	return plainText
}

func (sc *stringCondition) ToString() string {
	plainText := sc.StringProperty + " " + sc.StringComparator + " " + sc.StringValue
	return plainText
}




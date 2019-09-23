package cardobject

import "strconv"

func (cw *conditionWrapper) ToString() string {
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
	plainText := ic.Property + " " + ic.IntComparator + " " + strconv.Itoa(ic.IntValue)
	return plainText
}

func (sc *stringCondition) ToString() string {
	plainText := sc.Property + " " + sc.StringComparator + " " + sc.StringValue
	return plainText
}




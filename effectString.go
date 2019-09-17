package cardobject

import "strings"

func (e *effect) ToPlainText() string {
	var plainText string
	produce := e.Production
	manipulate := e.Manipulation
	move := e.ZoneChange
	if(len(produce) > 0) {
		plainText += "Produce " + strings.Join(produce, ",") + "."
	}
	if(manipulate != nil) {
		for i, m := range manipulate {
			if(i > 0) {
				plainText += " "
			}
    		plainText += m.ToPlainText()
		}
	}
	if(move != nil) {
		for i, m := range move {
			if(i > 0) {
				plainText += " "
			}
    		plainText += m.ToPlainText()
		}
	}
	return plainText
}
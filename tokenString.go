package cardobject

import "strconv"

func (t *token) toString() string {
	var plainText string
	name := t.Name
	tag := t.Tag

	if tag != nil {
		plainText += *tag + "-"
	}
	plainText += "token "
	if name != nil {
		plainText += "named " + *name
	}

	plainText += " with AbilitySpeed " + strconv.Itoa(int(t.AbilitySpeed))
	plainText += ", Health " + strconv.Itoa(int(t.Health))
	plainText += " and Attack " + strconv.Itoa(int(t.Attack))

	return plainText
}

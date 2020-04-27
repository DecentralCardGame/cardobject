package cardobject

import (
	"strconv"
	"strings"
)

func (r *ressources) toString() string {
	ressourceList := []string{}
	if r.Energy {
		ressourceList = append(ressourceList, strconv.FormatBool(r.Energy)+" Energy")
	}
	if r.Food {
		ressourceList = append(ressourceList, strconv.FormatBool(r.Food)+" Food")
	}
	if r.Lumber {
		ressourceList = append(ressourceList, strconv.FormatBool(r.Lumber)+" Lumber")
	}
	if r.Mana {
		ressourceList = append(ressourceList, strconv.FormatBool(r.Mana)+" Mana")
	}
	if r.Iron {
		ressourceList = append(ressourceList, strconv.FormatBool(r.Iron)+" Iron")
	}
	if r.Generic {
		ressourceList = append(ressourceList, strconv.FormatBool(r.Generic)+" Generic")
	}
	return strings.Join(ressourceList, ", ")
}

package cardobject

import "strconv"
import "strings"

func(r *ressources) toString() string {
  ressourceList := []string{}
  if(r.Energy){
    ressourceList = append(ressourceList, strconv.FormatBool(r.Energy) + " Energy")
  }
  if(r.Food){
    ressourceList = append(ressourceList, strconv.FormatBool(r.Food) + " Food")
  }
  if(r.Lumber){
    ressourceList = append(ressourceList, strconv.FormatBool(r.Lumber) + " Lumber")
  }
  if(r.Mana){
    ressourceList = append(ressourceList, strconv.FormatBool(r.Mana) + " Mana")
  }
  if(r.Metal){
    ressourceList = append(ressourceList, strconv.FormatBool(r.Metal) + " Metal")
  }
  if(r.Generic){
    ressourceList = append(ressourceList, strconv.FormatBool(r.Generic) + " Generic")
  }
  return strings.Join(ressourceList, ", ")
}
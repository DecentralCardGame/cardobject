package cardobject

import "strconv"
import "strings"

func(r *ressources) toString() string {
  ressourceList := []string{}
  if(r.Energy !=0){
    ressourceList = append(ressourceList, strconv.Itoa(int(r.Energy)) + " Energy")
  }
  if(r.Food !=0){
    ressourceList = append(ressourceList, strconv.Itoa(int(r.Food)) + " Food")
  }
  if(r.Lumber !=0){
    ressourceList = append(ressourceList, strconv.Itoa(int(r.Lumber)) + " Lumber")
  }
  if(r.Mana !=0){
    ressourceList = append(ressourceList, strconv.Itoa(int(r.Mana)) + " Mana")
  }
  if(r.Metal !=0){
    ressourceList = append(ressourceList, strconv.Itoa(int(r.Metal)) + " Metal")
  }
  if(r.Generic !=0){
    ressourceList = append(ressourceList, strconv.Itoa(int(r.Generic)) + " Generic")
  }
  return strings.Join(ressourceList, ", ")
}
package cardobject

import "testing"
import "io/ioutil"
import "encoding/json"
import "fmt"

func TestPlainText1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entity1Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Entity.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entity2Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Entity.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entity3Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Entity.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText4(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/field1Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Field.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText5(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/field2Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Field.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText6(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/field3Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Field.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText7(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hq1Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Headquarter.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText8(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hq2Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Headquarter.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText9(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hq3Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	abilities := card.Headquarter.Abilities
	for _, ability := range abilities {
    	fmt.Println(ability.ToPlainText())
	}
}
func TestPlainText10(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action1Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	effect := card.Action.Effect
	fmt.Println(effect.ToPlainText())
}
func TestPlainText11(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action2Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	effect := card.Action.Effect
	fmt.Println(effect.ToPlainText())
}
func TestPlainText12(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action3Test.json")
	cardJson := string(file)
	var card cardWrapper
    err := json.Unmarshal([]byte(cardJson), &card)
    if (err != nil) {
		t.Error(err)
	}
	effect := card.Action.Effect
	fmt.Println(effect.ToPlainText())
}
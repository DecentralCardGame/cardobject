package cardobject

import "testing"
import "fmt"
import "io/ioutil"
import "strings"

func TestCardSerialization1(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/entity2.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization2(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/actionTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization3(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/fieldTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization4(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/hqTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization5(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/entityTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}

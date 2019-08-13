package cardobject

import "testing"
import "fmt"
import "io/ioutil"
import "strings"

func TestCardSerialization2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/actionTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/fieldTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization4(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hqTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}
func TestCardSerialization5(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entityTest.json")
	input := string(file)
	output := ProcessCard(input)
 	//fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}

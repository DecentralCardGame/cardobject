package cardobject

import "testing"
import "io/ioutil"
import "strings"

func TestCardSerialization2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/actionTest.json")
	input := string(file)
	output, err := ProcessCard(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal")
	}
}
func TestCardSerialization3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/fieldTest.json")
	input := string(file)
	output, err := ProcessCard(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal")
	}
}
func TestCardSerialization4(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hqTest.json")
	input := string(file)
	output, err := ProcessCard(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal")
	}
}
func TestCardSerialization5(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entityTest.json")
	input := string(file)
	output, err := ProcessCard(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal")
	}
}

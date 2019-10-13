package cardobject

import "testing"
import "io/ioutil"
import "strings"

func TestCardSerializationAction1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action1Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationAction2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action2Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationAction3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/action3Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationField1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/field1Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationField2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/field2Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationField3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/field3Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationHq1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hq1Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationHq2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hq2Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationHq3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/hq3Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationEntity1(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entity1Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationEntity2(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entity2Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}
func TestCardSerializationEntity3(t *testing.T) {
	file, _ := ioutil.ReadFile("testJsons/entity3Test.json")
	input := string(file)
	output, err := FunctionalCardJson(input)
	if err != nil {
		t.Error(err)
	}
	if(strings.Compare(input, output) != 0){
		t.Errorf("In- and output are not equal: " + output)
	}
}

package cardobject

import "testing"
import "fmt"
import "io/ioutil"
import "strings"

func TestCardSerialization(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/entity2.json")
	input := string(file)
	output := ProcessCard(input)
 	fmt.Println(output)
 	fmt.Println(strings.Compare(input, output))
}

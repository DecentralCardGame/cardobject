package cardobject

import "testing"
import "fmt"
import "io/ioutil"

func TestCardSerialization(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/entity.json")
 	fmt.Println(ProcessCard(string(file)))
}

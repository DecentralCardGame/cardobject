package cardobject

import "testing"
import "fmt"

func TestCardSerialization(t *testing.T) {
	x := CreateAction()
	fmt.Println(SerializeCard(x))
}
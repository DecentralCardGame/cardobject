package cardobject

import "testing"
import "fmt"

func TestCardSerialization(t *testing.T) {
	x := createAction()
	fmt.Println(SerializeCard(x))
}
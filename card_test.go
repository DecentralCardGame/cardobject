package cardobject

import "testing"
import "fmt"

func TestCard(t *testing.T) {
	fmt.Println(createTestString())
}

func success(s string) {
	fmt.Println("[SUCCESS] " + s + " tested")
}

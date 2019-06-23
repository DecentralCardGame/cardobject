package cardobject

import "testing"
import "fmt"
import "io/ioutil"
import amino "github.com/tendermint/go-amino"

func TestCardSerialization(t *testing.T) {
	file, _ := ioutil.ReadFile("schema/entity.json")
 	fmt.Println(ProcessCard(string(file)))
 	var cardBuffer card

	cdc := amino.NewCodec()
	cdc.RegisterInterface((*card)(nil), nil)
	cdc.RegisterConcrete(action{}, "action", nil)
	cdc.RegisterConcrete(entity{}, "entity", nil)

	//json, _ := cdc.MarshalJSON(action{})

	cdc.MustUnmarshalJSON([]byte(file), &cardBuffer)

}

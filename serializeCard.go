package cardobject

import "fmt"
import "encoding/json"

func SerializeCard(c Card) string {
	bytes, err := json.Marshal(c)
    if err != nil {
        fmt.Println("Can't serialize", c)
        return "Can't serialize"
    }
    return string(bytes)
}

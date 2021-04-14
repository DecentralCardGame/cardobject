# Cardobject
The Cardobject describes a crowd-control card. A card is represented by a nested go-struct which contains all necessary data for a card including effects and abilities.
Addidtional functions provided allow to retrieve a build-instruction, an unmarshaling-method and the validation of a card.

# Usage
1. An instruction on how to build a card is provided in form of a json-schema.
2. A card-json can be unmarshaled and validated to a [Card](card.go#L7) by the [Unmarshal()](unmarshal.go#L5)-method.
3. The [Card](card.go#L7) provides a [Validate()](card.go#L14)-method which checks if the [Card](card.go#L7) is built correctly.

## Schema
1. The schema describes to a frontend how to build a keyworded card.
2. It is provided by the [Schema()](schema.go#L7)-function.

## Unmarshal and validate a card
1. The [Unmarshal()](unmarshal.go#L5)-method takes a keyworded card in form of a json and returns a [Card](card.go#L7) or an error.
2. This process involves validating the [Card](card.go#L7).
4. Example:
```golang
import "github.com/DecentralCardGame/cardobject/cardobject"

data, _ := ioutil.ReadFile("card.json")
  
card, err := cardobject.Unmarshal(data)
```

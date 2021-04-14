# Cardobject
The Cardobject describes a crowd-control card. A [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7) is represented by a nested go-struct which contains all necessary data  including effects and abilities.
Addidtional functions provided allow to retrieve a build-instruction, an unmarshaling-method and the validation of a card.
It also provides a shortened Version, the [KeywordedCard](https://github.com/DecentralCardGame/cardobject/tree/master/keywords/card.go#L12) which can be resolved to a full [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7).

# Usage
1. An instruction on how to build a card is provided in form of a json-schema.
2. A card-json can be unmarshaled and validated to a [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7).

## Schema
1. The schema that describes to a frontend how to build a card is provided by the [CardSchema()](schema.go#L17)-function.
2. The schema that describes to a frontend how to build a keyworded card is provided by the [KeywordedSchema()](schema.go#L10)-function.

## Unmarshal and validate a Card
1. The [UnmarshalRaw()](cardobject.go#L10)-method takes a [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7) in form of a json and returns a [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7) or an error.
2. This process involves validating the [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7).
3. Example:
```golang
import "github.com/DecentralCardGame/cardobject"

data, _ := ioutil.ReadFile("card.json")
  
card, err := cardobject.UnmarshalRaw(data)
```

## Unmarshal and validate a KeywordedCard
1. The [UnmarshalKeyworded()](cardobject.go#L20)-method takes a [KeywordedCard](https://github.com/DecentralCardGame/cardobject/tree/master/keywords/card.go#L12) in form of a json and returns a [Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject/card.go#L7) or an error.
2. This process involves resolving and validating the [KeywordedCard](https://github.com/DecentralCardGame/cardobject/tree/master/keywords/card.go#L12)
and validating the resulting [Card](https://github.com/DecentralCardGame/cardobject/tree/master/keywords/card.go#L12).
3. Example:
```golang
import "github.com/DecentralCardGame/cardobject"

data, _ := ioutil.ReadFile("KeywordedCard.json")
  
card, err := cardobject.UnmarshalKeyworded(data)
```

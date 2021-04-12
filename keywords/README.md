# Keywords
Keywords are shortcuts that allow to blackbox complex effects/ablities of a [cardobject.Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject) behind a simpler description.
Example:
> Arm:{Amount:X}

resolves to

>  "TargetEffect":{
      "EntityTargetEffect":{
         "EntitySelector":{
            "PlayerMode":"YOU",
            "CardMode":"ALL",
            "EntityZone":"FIELD"
         },
         "EntityManipulations":[
            {
               "EntityIntManipulation":{
                  "IntProperty":"HEALTH",
                  "IntOperator":"ADD",
                  "IntValue":{
                     "SimpleIntValue":2
                  }
               }
            },
            {
               "EntityIntManipulation":{
                  "IntProperty":"ATTACK",
                  "IntOperator":"ADD",
                  "IntValue":{
                     "SimpleIntValue":2
                  }
               }
            }
         ]
      }
  }

# Usage
1. An instruction on how to build a keyworded card is provided in form of a json-schema.
2. A keyworded card-json can be unmarshaled and validated to a [KeywordedCard](card.go#L12) by the [Unmarshal()](unmarshal.go#L10)-method.
3. The [KeywordedCard](card.go#L12) provides a [Resolve()](card.go#L13)-method which returns a [cardobject.Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject).

## Schema
1. The schema describes to a frontend how to build a keyworded card.
2. It is found [here](schema.json).

## Resolve keyworded card
1. The [Unmarshal()](unmarshal.go#L10)-method takes a keyworded card in form of a json and returns a [KeywordedCard](card.go#L12) or an error.
2. This process involves validating the keyworded card.
3. The [KeywordedCard](card.go#L12) then can be resolved to a [cardobject.Card](https://github.com/DecentralCardGame/cardobject/tree/master/cardobject).
4. Example:
```golang
import "github.com/DecentralCardGame/keywords"

data, _ := ioutil.ReadFile("keywordedCard.json")
  
keywordedCard, unmarshalErr := keywords.Unmarshal(data)
card, resolveErr := keywordedCard.Reslove()
```

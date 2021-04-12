# Keywords
Keywords are shortcuts that allow to blackbox complex effects/ablities of a [card](https://github.com/DecentralCardGame/cardobject) behind a simpler description.
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
2. A keyworded card can be validated and resolved to a [cardobject.Card](https://github.com/DecentralCardGame/cardobject) by the [Unmarshal](validator.go)-method.
3. The [cardobject.Card](https://github.com/DecentralCardGame/cardobject) is validated as well, so there is no need to use the [cardobject](https://github.com/DecentralCardGame/cardobject).

## Schema
1. The schema describes to a frontend how to build a keyworded card.
2. It is found [here](schema.json).

## Resolve keyworded card
1. The [Unmarshal](validator.go) takes a keyworded card in form of a json-string and returns a [cardobject.Card](https://github.com/DecentralCardGame/cardobject) or an error.
2. This process involves validating the keyworded card and resolving it to a [cardobject.Card](https://github.com/DecentralCardGame/cardobject) which is validated as well.
3. Example:
```golang
import "github.com/DecentralCardGame/keywords"

data, _ := ioutil.ReadFile("keywordedCard.json")
  
card, error := keywords.Unmarshal(data)
```

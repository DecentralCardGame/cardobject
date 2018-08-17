package card

import basics "github.com/DecentralCardGame/parser/basics"

type Card struct{
	//cost
	speedmodifier basics.Speedmodifier
	//tag basics.Tag <- array
	text basics.Text
}

type Entity struct{
	*Card
	//ability
	damage basics.Damage
	life basics.Life
}

type Location struct{
	*Card
	//ability
	life basics.Life
}

type Action struct{
	*Card
	//effect
}
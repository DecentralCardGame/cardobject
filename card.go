package cardobject

type Card struct{
	cost Cost
	speedmodifier Speedmodifier
	tag []Tag
	text Text
}

type Entity struct{
	*Card
	ability Ability
	damage Damage
	life Life
}

type Location struct{
	*Card
	ability Ability
	life Life
}

type Action struct{
	*Card
	//effect
}
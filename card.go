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

type Field struct{
	*Card
	storage []Ressource
	ability Ability
	life Life
	inhabitants []Entity
}

type Action struct{
	*Card
	effect Effect
}
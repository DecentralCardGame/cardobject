package cardobject

type Cost struct{
	ressourceCost []Ressource
	exhaustCost ExhaustCost
	zoneChangeCost []ZoneChange
}

type ExhaustCost struct{
	exhaustCostType string
}


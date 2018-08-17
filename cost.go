package cardobject

type Cost struct{
	ressourceCost []Ressource
	exhaustCost ExhaustCost
	//Zonechangecost
}

type ExhaustCost struct{
	exhaustCostType string
}


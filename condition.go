package cardobject

type Condition struct {}

type CardCondition interface {
	GetCardProperty() CardProperty
}

type PlayerCondition interface {

}

type cardCondition struct {
	Property CardProperty
}
func (cc *cardCondition) GetCardProperty() CardProperty{
	return cc.Property
}

type CardConditionInt struct{
	*cardCondition
	Property CardIntProperty
	Comparator Comparator
}

type CardConditionString struct{
	Property CardStringProperty
}
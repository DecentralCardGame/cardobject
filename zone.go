package cardobject

type Zone interface {
	IsZone() bool
}

type DynamicZone interface {
	Zone
	IsDynamic() bool
}

func DECK() *deck {
	return &deck{0}
}

const (
	DUSTPILE simpleZoneID = iota
	FIELD
	HAND
	QUEUE
)

const (
	EXILE protectedZoneID = iota
)

type simpleZoneID int

type protectedZoneID int

type deck struct {
	cardPosition int
} 


func (sz simpleZoneID) IsZone() bool {
	return true
}

func (pz protectedZoneID) IsZone() bool {
	return true
}

func (d deck) IsZone() bool{
	return true	
}

func (sz simpleZoneID) IsDynamic() bool {
	return true
}

func (d deck) IsDynamic() bool {
	return true
}



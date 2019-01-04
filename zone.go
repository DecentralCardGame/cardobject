package cardobject

type Zone interface {
	IsZone() bool
}

type DynamicZone interface {
	Zone
	IsDynamic() bool
}

func DECKRANGE(s int, e int) *deckRange {
	return &deckRange{s, e}
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

type deckRange struct {
	Start int
	End int
} 


func (sz simpleZoneID) IsZone() bool {
	return true
}

func (pz protectedZoneID) IsZone() bool {
	return true
}

func (d deckRange) IsZone() bool{
	return true	
}

func (sz simpleZoneID) IsDynamic() bool {
	return true
}

func (d deckRange) IsDynamic() bool {
	return true
}



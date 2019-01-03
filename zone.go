package cardobject

type Zone interface {
	IsZone() bool
}

type DynamicZone interface {
	Zone
	IsDynamic() bool
}

func NewDeckRange(s int, e int) *DeckRange {
	return &DeckRange{s, e}
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

type DeckRange struct {
	Start int
	End int
} 


func (sz simpleZoneID) IsZone() bool {
	return true
}

func (pz protectedZoneID) IsZone() bool {
	return true
}

func (d DeckRange) IsZone() bool{
	return true	
}

func (sz simpleZoneID) IsDynamic() bool {
	return true
}

func (d DeckRange) IsDynamic() bool {
	return true
}



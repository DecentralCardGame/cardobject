package cardobject

type protectedZoneID int
type moveZoneID int

const (
	DECK	 moveZoneID = iota
	DUSTPILE
	FIELD
	HAND
	EXILE	 protectedZoneID = iota
	QUEUE
)

type Zone interface {
	ZoneID()
}

type MoveZone interface {
	MoveZoneID()
}

type ProtectedZone interface {
	ProtectedZoneID()
}

func (z moveZoneID) ZoneID() {

}

func (z protectedZoneID) ZoneID() {

}

func (z moveZoneID) MoveZoneID() {

}

func (z protectedZoneID) ProtectedZoneID() {

}
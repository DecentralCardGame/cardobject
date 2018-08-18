package cardobject

import "strconv"
import "fmt"

type protectedZoneID int
type moveZoneID int

const (
	DECK	 moveZoneID = iota
	DUSTPILE
	FIELD
	HAND
	EXILE	 protectedZoneID = iota
	STACK
)

type Zone interface {
	ZoneID() int
}

type MoveZone interface {
	MoveZoneID() int
}

type ProtectedZone interface {
	ProtectedZoneID() int
}

func (z moveZoneID) ZoneID() int {
	return z.MoveZoneID()
}

func (z protectedZoneID) ZoneID() int {
	return z.ProtectedZoneID()
}

func (z moveZoneID) MoveZoneID() int {
	i, err := strconv.Atoi(string(z))
	if err != nil {

	}
	fmt.Println(z)
	return i
}

func (z protectedZoneID) ProtectedZoneID() int {
	i, err := strconv.Atoi(string(z))
	if err != nil {
		
	}
	fmt.Println(z)
	return i
}
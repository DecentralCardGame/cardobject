package cardobject

import "reflect"
import "errors"

type CardLocation interface {
	GetLocation() Zone
}

type DynamicCardLocation interface {
	GetDynamicLocation() DynamicZone
}

type Zone interface {
	IsZone() bool
}

type DynamicZone interface {
	IsZone() bool
	IsDynamic() bool
}

func NewCardLocation(z Zone) (CardLocation, error){
	szType := reflect.TypeOf(FIELD)
	pzType := reflect.TypeOf(EXILE)
	dzType := reflect.TypeOf(DECK{0})

	zType := reflect.TypeOf(z)

	switch zType {
		case szType: return &dynamicCardLocation{&cardLocation{z.(simpleZoneID)}}, nil
		case pzType: return &cardLocation{z.(protectedZoneID)}, nil
		case dzType: return &dynamicCardLocation{&cardLocation{z.(DECK)}}, nil
		default: return &cardLocation{}, errors.New("unknown Type")
	}
}

func NewDynamicCardLocation(z DynamicZone) (DynamicCardLocation, error){
	l, err :=NewCardLocation(z)
	return l.(DynamicCardLocation), err
}

type cardLocation struct {
	zone Zone 
}

type dynamicCardLocation struct {
	*cardLocation
}

func (cl *cardLocation) GetLocation() Zone {
	return cl.zone
}

func (dcl *dynamicCardLocation) GetDynamicLocation() DynamicZone {
	return dcl.zone.(DynamicZone)
}

type simpleZoneID int
type protectedZoneID int

const (
	DUSTPILE simpleZoneID = iota
	FIELD
	HAND
	QUEUE
)

const (
	EXILE protectedZoneID = iota
)

type DECK struct {
	cardPosition int
} 

func (sz simpleZoneID) IsZone() bool {
	return true
}

func (pz protectedZoneID) IsZone() bool {
	return true
}

func (d DECK) IsZone() bool{
	return true	
}



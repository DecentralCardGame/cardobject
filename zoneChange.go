package cardobject

type ZoneChange interface {
	GetSelection() CardSelectorSafeCond
	GetDestination() Zone
}

func NewZoneChange(s CardSelectorSafeCond, z Zone) ZoneChange {
	return &zoneChange{s, z}
}


type zoneChange struct {
	Selection CardSelectorSafeCond
	Destination Zone
}


func (zc *zoneChange) GetSelection() CardSelectorSafeCond {
	return zc.Selection
}

func (zc *zoneChange) GetDestination() Zone {
	return zc.Destination
}
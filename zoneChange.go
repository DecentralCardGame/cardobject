package cardobject

type ZoneChange interface {
	GetSelection() CardSelectorSafeCond
	GetDestination() Zone
}

func NewZoneChange(s CardSelectorSafeCond, z Zone) ZoneChange {
	return &zoneChange{s, z}
}


type zoneChange struct {
	selection CardSelectorSafeCond
	destination Zone
}


func (zc *zoneChange) GetSelection() CardSelectorSafeCond {
	return zc.selection
}

func (zc *zoneChange) GetDestination() Zone {
	return zc.destination
}
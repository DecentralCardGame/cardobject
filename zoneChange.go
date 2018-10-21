package cardobject

type ZoneChange interface {
	GetSelection() SafeSelector
	GetDestination() Zone
}

func NewZoneChange(s SafeSelector, z Zone) ZoneChange {
	return &zoneChange{s, z}
}


type zoneChange struct {
	selection SafeSelector
	destination Zone
}


func (zc *zoneChange) GetSelection() SafeSelector {
	return zc.selection
}

func (zc *zoneChange) GetDestination() Zone {
	return zc.destination
}
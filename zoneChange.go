package cardobject

func NewZoneChange(s CardSelectorSafeCond, z Zone) *ZoneChange {
	return &ZoneChange{s, z}
}

type ZoneChange struct {
	selection CardSelectorSafeCond
	destination Zone
}


func (zc *ZoneChange) GetSelection() CardSelectorSafeCond {
	return zc.selection
}

func (zc *ZoneChange) GetDestination() Zone {
	return zc.destination
}
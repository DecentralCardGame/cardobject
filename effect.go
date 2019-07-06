package cardobject

type effect struct {
	Production []string
	Manipulation []manipulationWrapper
	ZoneChange []zoneChange
}
package bucontracts

type FleetContactBO struct {
	Id        uint
	FleetId   uint
	ContactId uint
	Fleet     FleetBO
	Contact   ContactBO
}

package bucontracts

type FleetAddressBO struct {
	Id        uint
	FleetId   uint
	AddressId uint
	Fleet     FleetBO
	Address   AddressBO
}

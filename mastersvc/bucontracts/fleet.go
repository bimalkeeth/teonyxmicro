package bucontracts

import "time"

type FleetBO struct {
	Id                   uint
	UpdatedAt            time.Time
	FleetId              string
	Name                 string
	SurName              string
	OtherName            string
	DateRegistered       time.Time
	RegistrationDuration int
	FleetContactId       uint
	FleetLocationId      uint

	FleetContacts []ContactBO
	Address       []AddressBO
}

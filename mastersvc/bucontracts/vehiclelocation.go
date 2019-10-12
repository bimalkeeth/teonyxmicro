package bucontracts

import "time"

type VehicleAddressBO struct {
	Id        uint
	AddressId uint
	VehicleId uint
	UpdateAt  time.Time
	Address   AddressBO
}

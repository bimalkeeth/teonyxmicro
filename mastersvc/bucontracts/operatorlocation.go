package bucontracts

type OperatorLocationBO struct {
	Id         uint
	AddressId  uint
	OperatorId uint
	Address    AddressBO
	Operator   *VehicleOperatorBO
}

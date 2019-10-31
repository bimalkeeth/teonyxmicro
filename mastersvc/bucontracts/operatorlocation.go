package bucontracts

type OperatorLocationBO struct {
	Id         uint
	AddressId  uint
	OperatorId uint
	Primary    bool
	Address    AddressBO
	Operator   *VehicleOperatorBO
}

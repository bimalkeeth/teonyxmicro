package bucontracts

type VehicleOperatorBoundBO struct {
	Id         uint
	OperatorId uint
	VehicleId  uint
	Active     bool
	Operator   *VehicleOperatorBO
	Vehicle    *VehicleBO
}

package business

import bu "teonyxmicro/mastersvc/bucontracts"

type IVehicleOperatorBound interface {
	CreateVehicleOpBound(op bu.VehicleOperatorBoundBO) (uint, error)
	UpdateVehicleOpBound(op bu.VehicleOperatorBoundBO) (bool, error)
	DeleteVehicleOpBound(id uint) (bool, error)
}

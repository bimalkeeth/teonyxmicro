package vehicles

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

func (v *VehicleManager) CreateVehicleReg(bo bu.VehicleTrackRegBO) (uint, error) {
	vr := vehicleFac.New(bs.CVhRegistration).(bs.VehicleReg)
	result, err := vr.CreateVehicleReg(bo)
	return result, err
}

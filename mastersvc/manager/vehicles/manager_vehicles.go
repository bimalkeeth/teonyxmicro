package vehicles

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

//----------------------------------------
//Create Vehicle
//----------------------------------------
func (v *VehicleManager) CreateVehicle(vehicle bu.VehicleBO) (uint, error) {
	veh := vehicleFac.New(bs.CVehicle).(bs.Vehicle)
	vehicleFac.Conn.Begin()
	res, err := veh.CreateVehicle(vehicle)

	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

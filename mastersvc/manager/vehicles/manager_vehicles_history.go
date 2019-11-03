package vehicles

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

//-----------------------------------------------
//Create Vehicle History
//-----------------------------------------------
func (v *VehicleManager) CreateVehicleHistory(history bu.VehicleHistoryBO) (uint, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(bs.VehicleHistory)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleHistory(history)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------------
//Update Vehicle History
//-----------------------------------------------
func (v *VehicleManager) UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(bs.VehicleHistory)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleHistory(history)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//---------------------------------------------
//Delete Vehicle History
//---------------------------------------------
func (v *VehicleManager) DeleteVehicleHistory(id uint) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(bs.VehicleHistory)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleHistory(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

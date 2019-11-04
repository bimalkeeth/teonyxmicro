package vehicles

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

//---------------------------------------------
//Create vehicle make
//---------------------------------------------
func (v *VehicleManager) CreateVehicleMake(bo bu.VehicleMakeBO) (uint, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(bs.VehicleMake)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleMake(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//--------------------------------------------
//Update vehicle make
//--------------------------------------------
func (v *VehicleManager) UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(bs.VehicleMake)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleMake(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//--------------------------------------------
//Delete Vehicle make
//--------------------------------------------
func (v *VehicleManager) DeleteVehicleMake(id uint) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(bs.VehicleMake)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleMake(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//------------------------------------------
//Get all vehicle make
//------------------------------------------
func (v *VehicleManager) GetAllVehicleMake() ([]bu.VehicleMakeBO, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(bs.VehicleMake)
	res, err := vh.GetAllVehicleMake()
	return res, err
}

//------------------------------------------
//Get  vehicle make by id
//------------------------------------------
func (v *VehicleManager) GetVehicleMakeById(id uint) (bu.VehicleMakeBO, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(bs.VehicleMake)
	res, err := vh.GetVehicleMakeById(id)
	return res, err
}

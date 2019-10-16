package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type IVehicle interface {
	CreateVehicle(vehicle bu.VehicleBO) (uint, error)
	UpdateVehicle(vehicle bu.VehicleBO) (bool, error)
	DeleteVehicle(vehicleId uint) (bool, error)
	GetVehicleById(vehicleId uint) (bu.VehicleBO, error)
	GetVehicleByCountryId(countryId uint) ([]bu.VehicleBO, error)
	GetVehiclesByFleetId(fleetId uint) ([]bu.VehicleBO, error)
}
type Vehicle struct{ Db *gorm.DB }

func NewVehicle(db *gorm.DB) Vehicle {
	return Vehicle{Db: db}
}

//----------------------------------------------------
// Create Vehicle
//----------------------------------------------------
func (v *Vehicle) CreateVehicle(vehicle bu.VehicleBO) (uint, error) {
	vehicleTable := ent.TableVehicle{MakeId: vehicle.MakeId,
		FleetId:      vehicle.FleetId,
		ModelId:      vehicle.ModelId,
		Registration: vehicle.Registration,
		StatusId:     vehicle.StatusId}
	v.Db.Create(&vehicleTable)
	return vehicleTable.ID, nil
}

//----------------------------------------------------
// Update Vehicle
//---------------------------------------------------
func (v *Vehicle) UpdateVehicle(vehicle bu.VehicleBO) (bool, error) {
	vehicleTable := ent.TableVehicle{}
	v.Db.First(&vehicleTable, vehicle.Id)
	if vehicleTable.ID == 0 {
		return false, errors.New("vehicle not found for update")
	}
	vehicleTable.MakeId = vehicle.MakeId
	vehicleTable.StatusId = vehicle.StatusId
	vehicleTable.Registration = vehicle.Registration
	vehicleTable.ModelId = vehicle.ModelId
	vehicleTable.FleetId = vehicle.FleetId
	v.Db.Save(&vehicleTable)
	return true, nil
}

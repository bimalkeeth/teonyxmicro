package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type IVehicleStatus interface {
	CreateVehicleStatus(bo bu.VehicleStatusBO) (uint, error)
	UpdateVehicleStatus(bo bu.VehicleStatusBO) (bool, error)
	DeleteVehicleStatus(id uint) (bool, error)
	GetAllVehicleStatus() ([]bu.VehicleStatusBO, error)
}

type VehicleStatus struct {
	Db *gorm.DB
}

func NewVhStatus(db *gorm.DB) *VehicleStatus {
	return &VehicleStatus{Db: db}
}

//-----------------------------------------------
//Create vehicle status
//-----------------------------------------------
func (v *VehicleStatus) CreateVehicleStatus(bo bu.VehicleStatusBO) (uint, error) {
	vhStatus := ent.TableVehicleStatus{
		StatusName: bo.StatusName,
		StatusType: bo.StatusType,
	}
	v.Db.Create(&vhStatus)
	return vhStatus.ID, nil
}

//-----------------------------------------------
//Update vehicle status
//-----------------------------------------------
func (v *VehicleStatus) UpdateVehicleStatus(bo bu.VehicleStatusBO) (bool, error) {
	vhStatus := ent.TableVehicleStatus{}
	v.Db.First(&vhStatus, bo.Id)
	if vhStatus.ID == 0 {
		return false, errors.New("record not found")
	}

	vhStatus.StatusType = bo.StatusType
	vhStatus.StatusName = bo.StatusName
	v.Db.Save(&vhStatus)
	return true, nil
}

//-----------------------------------------------
//Delete vehicle status
//-----------------------------------------------
func (v *VehicleStatus) DeleteVehicleStatus(id uint) (bool, error) {
	vhStatus := ent.TableVehicleStatus{}
	v.Db.First(&vhStatus, id)
	if vhStatus.ID == 0 {
		return false, errors.New("record not found")
	}
	v.Db.Delete(&vhStatus)
	return true, nil
}

//-----------------------------------------------
//Get all vehicle status
//-----------------------------------------------
func (v *VehicleStatus) GetAllVehicleStatus() ([]bu.VehicleStatusBO, error) {
	var vhStatus []ent.TableVehicleStatus
	var vhStatusResult []bu.VehicleStatusBO
	v.Db.Find(&vhStatus)

	for _, item := range vhStatus {
		vhStatusResult = append(vhStatusResult, bu.VehicleStatusBO{
			StatusName: item.StatusName,
			StatusType: item.StatusType,
			UpdatedAt:  item.UpdatedAt,
		})
	}
	return vhStatusResult, nil
}

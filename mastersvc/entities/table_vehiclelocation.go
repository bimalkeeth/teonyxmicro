package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleLocation struct {
	gorm.Model
	AddressId uint `gorm:"column:addressid;not_null"`
	VehicleId uint `gorm:"column:vehicleid;not_null"`
	Address   *TableAddress
	Vehicle   *TableVehicle
}

func (t TableVehicleLocation) TableName() string {
	return "table_vehiclelocation"
}

func (t TableVehicleLocation) Validate(db *gorm.DB) {
	if t.AddressId == 0 {
		_ = db.AddError(errors.New("address should contain value"))
	}
	if t.VehicleId == 0 {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
}

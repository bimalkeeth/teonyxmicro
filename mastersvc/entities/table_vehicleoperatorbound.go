package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleOperatorBound struct {
	gorm.Model
	OPeratorId uint `gorm:"column:operatorid;not_null"`
	VehicleId  uint `gorm:"column:operatorid;not_null"`
	Operator   *TableVehicleOperators
	Vehicle    *TableVehicle
}

func (t TableVehicleOperatorBound) TableName() string {
	return "table_vehicleoperatorbound"
}

func (t TableVehicleOperatorBound) Validate(db *gorm.DB) {

	if t.OPeratorId == 0 {
		_ = db.AddError(errors.New("operator should contain value"))
	}
	if t.VehicleId == 0 {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
}

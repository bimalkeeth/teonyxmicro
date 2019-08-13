package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleOperators struct {
	gorm.Model
	Name       string `gorm:"column:name;not_null"`
	SurName    string `gorm:"column:surname;not_null"`
	DrivingLic string `gorm:"column:drivinglic;not_null"`
}

func (t TableVehicleOperators) TableName() string {
	return "table_vehicleoperators"
}

func (t TableVehicleOperators) Validate(db *gorm.DB) {

	if len(t.Name) == 0 {
		_ = db.AddError(errors.New("name should contain value"))
	}
	if len(t.SurName) == 0 {
		_ = db.AddError(errors.New("surname should contain value"))
	}
	if len(t.DrivingLic) == 0 {
		_ = db.AddError(errors.New("driving license should contain value"))
	}
}

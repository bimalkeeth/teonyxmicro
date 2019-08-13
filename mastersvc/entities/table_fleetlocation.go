package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableFleetLocation struct {
	gorm.Model
	FleetId   uint          `gorm:"column:fleetid;not_null;unique_index:fleetlocation_fleetaddress_uindex"`
	AddressId uint          `gorm:"column:addressid;not_null;unique_index:fleetlocation_fleetaddress_uindex"`
	Fleet     *TableFleet   `gorm:"foreignkey:fleetid"`
	Address   *TableAddress `gorm:"foreignkey:addressid"`
}

func (t TableFleetLocation) TableName() string {
	return "table_fleetlocation"
}

func (t TableFleetLocation) Validate(db *gorm.DB) {

	if t.FleetId == 0 {

		_ = db.AddError(errors.New("fleet should contain value"))
	}
	if t.AddressId == 0 {

		_ = db.AddError(errors.New("address should contain value"))
	}
}

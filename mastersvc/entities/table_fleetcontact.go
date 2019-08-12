package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableFleetContact struct {
	gorm.Model
	ContactId uint `gorm:"column:contactid;not_null"`
	FleetId   uint `gorm:"column:fleetid;not_null"`
	Contact   *TableContact
	Fleet     *TableFleet
}

func (t TableFleetContact) TableName() string {
	return "table_fleetcontact"
}

func (t TableFleetContact) Validate(db *gorm.DB) {
	if t.ContactId == 0 {
		_ = db.AddError(errors.New("contact should contain value"))
	}
	if t.FleetId == 0 {
		_ = db.AddError(errors.New("fleet should contain value"))
	}
}

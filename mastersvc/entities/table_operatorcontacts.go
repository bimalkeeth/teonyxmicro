package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableOPeratorContacts struct {
	gorm.Model
	ContactId  uint `gorm:"column:contactid;not_null"`
	OPeratorId uint `gorm:"column:operatorid;not_null"`
	Contact    *TableContact
	Operator   *TableVehicleOperators
}

func (t TableOPeratorContacts) TableName() string {
	return "table_operatorcontacts"
}

func (t TableOPeratorContacts) Validate(db *gorm.DB) {

	if t.ContactId == 0 {

		_ = db.AddError(errors.New("contact should contain value"))
	}
	if t.OPeratorId == 0 {

		_ = db.AddError(errors.New("operator should contain value"))
	}
}

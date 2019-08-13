package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableOPeratorContacts struct {
	gorm.Model
	ContactId  uint                   `gorm:"column:contactid;not_null;unique_index:operatorcontacts_contoperator_uindex"`
	OperatorId uint                   `gorm:"column:operatorid;not_null;unique_index:operatorcontacts_contoperator_uindex"`
	Contact    *TableContact          `gorm:"foreignkey:contactid"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
}

func (t TableOPeratorContacts) TableName() string {
	return "table_operatorcontacts"
}

func (t TableOPeratorContacts) Validate(db *gorm.DB) {

	if t.ContactId == 0 {

		_ = db.AddError(errors.New("contact should contain value"))
	}
	if t.OperatorId == 0 {

		_ = db.AddError(errors.New("operator should contain value"))
	}
}

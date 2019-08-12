package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TableFleet struct {
	gorm.Model
	Name                 string    `gorm:"column:name;not_null"`
	SurName              string    `gorm:"column:surname"`
	OtherName            string    `gorm:"column:othernames"`
	DateRegistered       time.Time `gorm:"column:dateregistered;not_null"`
	RegistrationDuration int       `gorm:"column:regisrationduration;not_null"`
}

func (t TableFleet) TableName() string {
	return "table_fleet"
}

func (t TableFleet) Validate(db *gorm.DB) {

	if len(t.Name) == 0 {

		_ = db.AddError(errors.New("name should contain value"))
	}
	if t.RegistrationDuration == 0 {

		_ = db.AddError(errors.New("registration duration should contain value"))
	}
	if t.DateRegistered.IsZero() {

		_ = db.AddError(errors.New("registration date should contain value"))
	}
}

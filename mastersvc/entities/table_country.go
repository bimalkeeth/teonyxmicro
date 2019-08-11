package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableCountry struct {
	gorm.Model
	CountryName string `gorm:"column:countryname;not_null"`
	RegionId    uint   `gorm:"column:regionid;not_null"`
	Region      *TableRegion
}

func (t TableCountry) TableName() string {
	return "table_contacttype"
}

func (t TableCountry) Validate(db *gorm.DB) {

	if len(t.CountryName) == 0 {
		_ = db.AddError(errors.New("country name should contain value"))
	}
	if t.RegionId == 0 {
		_ = db.AddError(errors.New("region should contain value"))
	}
}

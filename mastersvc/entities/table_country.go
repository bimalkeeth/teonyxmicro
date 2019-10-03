package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableCountry struct {
	gorm.Model
	CountryName string        `gorm:"column:countryname;not_null;unique_index:country_countryname_uindex"`
	RegionId    uint          `gorm:"column:regionid;not_null"`
	Region      *TableRegion  `gorm:"foreignkey:regionid"`
	States      []*TableState `gorm:"foreignkey:CountryId;association_foreignkey:id"`
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

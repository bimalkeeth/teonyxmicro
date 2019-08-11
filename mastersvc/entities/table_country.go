package entities

import "github.com/jinzhu/gorm"

type TableCountry struct {
	gorm.Model
	CountryName string `gorm:"column:countryname;not_null"`
	RegionId    uint   `gorm:"column:regionid;not_null"`
	Region      *TableRegion
}

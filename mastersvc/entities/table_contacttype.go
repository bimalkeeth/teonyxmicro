package entities

import "github.com/jinzhu/gorm"

type TableAddress struct {
	gorm.Model
	Address       string `gorm:"column:address;not_null"`
	Street        string `gorm:"column:street:not_null"`
	Suberb        string `gorm:"column:suberb"`
	StateId       uint   `gorm:"column:stateid"`
	CountryId     uint   `gorm:"column:countryid"`
	AddressTypeId uint   `gorm:"column:addresstypeid;not_null"`
	Location      string `gorm:"column:location;not_null"`
}

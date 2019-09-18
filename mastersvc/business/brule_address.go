package business

import (
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
)

type IAddress interface {
	CreateAddress(db *gorm.DB, address bu.AddressBO) (bool, error)
	UpdateAddress(db *gorm.DB, address bu.AddressBO) (bool, error)
	DeleteAddress(db *gorm.DB, id uint) (bool, error)
	GetAddressById(db *gorm.DB, id uint) (bu.AddressBO, error)
	GetAddressByName(db *gorm.DB, name string) ([]bu.AddressBO, error)
}

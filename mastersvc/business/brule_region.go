package business

import (
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
)

type Region interface {
	CreateRegion(db *gorm.DB, bo bu.RegionBO)
	UpdateRegion(db *gorm.DB, bo bu.RegionBO)
	DeleteRegion(db *gorm.DB, Id uint)
}

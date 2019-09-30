package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	"teonyxmicro/mastersvc/entities"
)

//-----------------------------------------------
// Interface for region management
//-----------------------------------------------
type IRegion interface {
	CreateRegion(db *gorm.DB, bo bu.RegionBO) (bool, error)
	UpdateRegion(db *gorm.DB, bo bu.RegionBO) (bool, error)
	DeleteRegion(db *gorm.DB, id uint) (bool, error)
	GetAllRegion(db *gorm.DB) ([]bu.RegionBO, error)
	GetRegionById(db *gorm.DB, id uint) (bu.RegionBO, error)
	GetRegionByName(db *gorm.DB, name string) (bu.RegionBO, error)
}
type Region struct{}

func NewRegion() *Region { return &Region{} }

//------------------------------------------------
//Create region for the given data
//------------------------------------------------
func (r *Region) CreateRegion(db *gorm.DB, bo bu.RegionBO) (bool, error) {

	db.Create(&entities.TableRegion{Region: bo.Region, RegionName: bo.RegionName})
	return true, nil
}

//------------------------------------------------
//Update the region given
//------------------------------------------------
func (r *Region) UpdateRegion(db *gorm.DB, bo bu.RegionBO) (bool, error) {

	region := entities.TableRegion{}
	db.First(&region, bo.Id)
	if region.ID == 0 {
		return false, errors.New("no record found for region")
	}
	region.Region = bo.Region
	region.RegionName = bo.Region
	db.Save(&region)
	return true, nil
}

//------------------------------------------------
// Delete region by Id
//------------------------------------------------
func (r *Region) DeleteRegion(db *gorm.DB, id uint) (bool, error) {

	found := entities.TableRegion{}
	db.First(&found, id)
	if found.ID == 0 {
		return false, errors.New("contact type not found")
	}
	db.Delete(&found)
	return true, nil
}

//------------------------------------------------
//Get al region
//------------------------------------------------
func (r *Region) GetAllRegion(db *gorm.DB) ([]bu.RegionBO, error) {
	var regions []entities.TableRegion
	var result []bu.RegionBO

	db.Find(&regions)
	for _, item := range regions {
		result = append(result, bu.RegionBO{Region: item.Region, RegionName: item.RegionName, Id: item.ID})
	}
	return result, nil
}

//-----------------------------------------------
// Get region by Id
//-----------------------------------------------
func (r *Region) GetRegionById(db *gorm.DB, id uint) (bu.RegionBO, error) {
	region := &entities.TableRegion{}
	db.First(&region, id)

	result := bu.RegionBO{}
	if region.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.RegionBO{Region: region.Region, RegionName: region.RegionName, Id: region.ID}, nil
}

//-----------------------------------------------
// Get region by name
//-----------------------------------------------
func (r *Region) GetRegionByName(db *gorm.DB, name string) (bu.RegionBO, error) {
	region := entities.TableRegion{}
	db.Where(&entities.TableRegion{RegionName: name}).First(&region)
	if region.ID == 0 {
		return bu.RegionBO{}, errors.New("record not found")
	}
	return bu.RegionBO{Region: region.Region, RegionName: region.RegionName, Id: region.ID}, nil
}

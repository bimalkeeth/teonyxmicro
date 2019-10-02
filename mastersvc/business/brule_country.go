package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type ICountry interface {
	CreateCountry(db *gorm.DB, bo bu.CountryBO) (bool, error)
	UpdateCountry(db *gorm.DB, bo bu.CountryBO) (bool, error)
	DeleteCountry(db *gorm.DB, Id uint) (bool, error)
	GetCountryById(db *gorm.DB, id uint) (bu.CountryBO, error)
	GetAllCountries(db *gorm.DB) ([]bu.CountryBO, error)
}

type Country struct{}

func NewCountry() *Country { return &Country{} }

//--------------------------------------------
//Create country
//--------------------------------------------
func (c *Country) CreateCountry(db *gorm.DB, bo bu.CountryBO) (bool, error) {
	db.Create(&ent.TableCountry{CountryName: bo.CountryName,
		RegionId: bo.RegionId,
	})
	return true, nil
}

//---------------------------------------------
//Update country
//---------------------------------------------
func (c *Country) UpdateCountry(db *gorm.DB, bo bu.CountryBO) (bool, error) {
	country := ent.TableCountry{}
	db.First(&country, bo.Id)
	if country.ID == 0 {
		return false, errors.New("country not found")
	}
	country.RegionId = bo.RegionId
	country.CountryName = bo.CountryName
	db.Save(&country)
	return true, nil
}

//----------------------------------------------
//Delete Country
//----------------------------------------------
func (c *Country) DeleteCountry(db *gorm.DB, Id uint) (bool, error) {
	country := ent.TableCountry{}
	db.First(&country, Id)
	if country.ID == 0 {
		return false, errors.New("country not found")
	}
	db.Delete(&country)
	return true, nil
}

//------------------------------------------------
//Get Country by Id
//------------------------------------------------
func (c *Country) GetCountryById(db *gorm.DB, id uint) (bu.CountryBO, error) {

	country := ent.TableCountry{}
	db.First(&country, id).Related(&ent.TableRegion{})
	if country.ID == 0 {
		return bu.CountryBO{}, errors.New("country not found")
	}
	return bu.CountryBO{Id: country.ID,
		CountryName: country.CountryName,
		RegionId:    country.RegionId,
		Region: bu.RegionBO{
			Id:         country.Region.ID,
			Region:     country.Region.Region,
			RegionName: country.Region.RegionName,
		},
		UpdatedAt: country.UpdatedAt,
	}, nil
}

//-----------------------------------------------
//Get country
//-----------------------------------------------
func (c *Country) GetAllCountries(db *gorm.DB) ([]bu.CountryBO, error) {
	var result []bu.CountryBO
	var countries []ent.TableCountry
	db.Find(&countries)

	for _, item := range countries {
		result = append(result, bu.CountryBO{Id: item.ID,
			CountryName: item.CountryName,
			RegionId:    item.RegionId,
			Region: bu.RegionBO{
				Id:         item.Region.ID,
				Region:     item.Region.Region,
				RegionName: item.Region.RegionName,
			},
			UpdatedAt: item.UpdatedAt,
		})
	}
	return result, nil
}

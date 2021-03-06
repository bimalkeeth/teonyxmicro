package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type ICountry interface {
	CreateCountry(bo bu.CountryBO) (uint, error)
	UpdateCountry(bo bu.CountryBO) (bool, error)
	DeleteCountry(Id uint) (bool, error)
	GetCountryById(id uint) (bu.CountryBO, error)
	GetAllCountries() ([]bu.CountryBO, error)
}

type Country struct{ Db *gorm.DB }

func NewCountry(db *gorm.DB) *Country { return &Country{Db: db} }

//--------------------------------------------
//Create country
//--------------------------------------------
func (c *Country) CreateCountry(bo bu.CountryBO) (uint, error) {
	country := ent.TableCountry{CountryName: bo.CountryName,
		RegionId: bo.RegionId,
	}
	c.Db.Create(&country)
	return country.ID, nil
}

//---------------------------------------------
//Update country
//---------------------------------------------
func (c *Country) UpdateCountry(bo bu.CountryBO) (bool, error) {
	country := ent.TableCountry{}
	c.Db.First(&country, bo.Id)
	if country.ID == 0 {
		return false, errors.New("country not found")
	}
	country.RegionId = bo.RegionId
	country.CountryName = bo.CountryName
	c.Db.Save(&country)
	return true, nil
}

//----------------------------------------------
//Delete Country
//----------------------------------------------
func (c *Country) DeleteCountry(Id uint) (bool, error) {
	country := ent.TableCountry{}
	c.Db.First(&country, Id)
	if country.ID == 0 {
		return false, errors.New("country not found")
	}
	c.Db.Delete(&country)
	return true, nil
}

//------------------------------------------------
//Get Country by Id
//------------------------------------------------
func (c *Country) GetCountryById(id uint) (bu.CountryBO, error) {

	country := ent.TableCountry{}
	c.Db.Preload("Region").First(&country, id)
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
func (c *Country) GetAllCountries() ([]bu.CountryBO, error) {
	var result []bu.CountryBO
	var countries []ent.TableCountry
	c.Db.Preload("Region").Find(&countries)

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

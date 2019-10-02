package business

import (
	"errors"
	"github.com/jinzhu/gorm"
)
import bu "teonyxmicro/mastersvc/bucontracts"
import en "teonyxmicro/mastersvc/entities"

type IFleet interface {
	CreateFleet(db *gorm.DB, bo bu.FleetBO) (bu.FleetBO, error)
	UpdateFleet(db *gorm.DB, bo bu.FleetBO) (bool, error)
	DeleteFleet(db *gorm.DB, id uint) (bool, error)
	GetFleetById(db gorm.DB, id uint) (bu.FleetBO, error)
}

type Fleet struct{}

func NewFleet() Fleet {
	return Fleet{}
}

//--------------------------------------------------------
//Create Fleet
//--------------------------------------------------------
func (f *Fleet) CreateFleet(db *gorm.DB, bo bu.FleetBO) (bu.FleetBO, error) {

	fleet := &en.TableFleet{
		FleetId:              bo.FleetId,
		Name:                 bo.Name,
		SurName:              bo.SurName,
		OtherName:            bo.OtherName,
		DateRegistered:       bo.DateRegistered,
		RegistrationDuration: bo.RegistrationDuration,
	}
	db.Create(fleet)
	return bo, nil
}

//----------------------------------------------------------
//Update Fleet
//----------------------------------------------------------
func (f *Fleet) UpdateFleet(db *gorm.DB, bo bu.FleetBO) (bool, error) {

	fleet := &en.TableFleet{}
	db.First(fleet, bo.Id)
	if fleet.ID == 0 {
		return false, errors.New("fleet can not be found")
	}

	fleet.SurName = bo.SurName
	fleet.Name = bo.Name
	fleet.FleetId = bo.FleetId
	fleet.OtherName = bo.OtherName
	fleet.DateRegistered = bo.DateRegistered
	fleet.RegistrationDuration = bo.RegistrationDuration

	db.Save(fleet)

	return true, nil
}

//----------------------------------------------------------
//Update Fleet
//----------------------------------------------------------
func (f *Fleet) DeleteFleet(db *gorm.DB, id uint) (bool, error) {
	fleet := &en.TableFleet{}
	db.First(fleet, id)
	if fleet.ID == 0 {
		return false, errors.New("fleet can not be found")
	}
	db.Delete(fleet)
	return true, nil
}

//----------------------------------------------------------
//Get Fleet Id
//----------------------------------------------------------
func (f *Fleet) GetFleetById(db gorm.DB, id uint) (bu.FleetBO, error) {
	fleet := &en.TableFleet{}
	db.Preload("FleetContacts").Preload("FleetLocations").First(fleet, id)
	if fleet.ID == 0 {
		return bu.FleetBO{}, errors.New("fleet can not be found")
	}

	result := bu.FleetBO{}
	result.Id = fleet.ID
	result.RegistrationDuration = fleet.RegistrationDuration
	result.DateRegistered = fleet.DateRegistered
	result.OtherName = fleet.OtherName
	result.FleetId = fleet.FleetId
	result.SurName = fleet.SurName
	result.Name = fleet.Name

	var contacts []bu.ContactBO

	var contact = fleet.FleetContacts

	for _, item := range *contact {
		contacts = append(contacts, bu.ContactBO{
			Id:            item.Contact.ID,
			Contact:       item.Contact.Contact,
			ContactTypeId: item.Contact.ContactTypeId,
		})
	}

}

package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	en "teonyxmicro/mastersvc/entities"
)

type IContact interface {
	CreateContact(db *gorm.DB, con bu.ContactBO) (bool, error)
	UpdateContact(db *gorm.DB, con bu.ContactBO) (bool, error)
	DeleteContact(db *gorm.DB, Id uint) (bool, error)
	ContactById(db *gorm.DB, Id uint) (bu.ContactBO, error)
}

type Contact struct{}

func NewContact() *Contact { return &Contact{} }

//--------------------------------------------
//Create Contact
//---------------------------------------------
func (c *Contact) CreateContact(db *gorm.DB, con bu.ContactBO) (bool, error) {

	db.Create(&en.TableContact{Contact: con.Contact, ContactTypeId: con.ContactTypeId})
	return true, nil
}

//--------------------------------------------
//Update Contact
//---------------------------------------------
func (c *Contact) UpdateContact(db *gorm.DB, con bu.ContactBO) (bool, error) {
	contact := en.TableContact{}
	db.First(&contact, con.Id)
	if contact.ID == 0 {
		return false, errors.New("contact id cannot be found")
	}
	contact.ContactTypeId = con.ContactTypeId
	contact.Contact = con.Contact
	db.Save(&contact)
	return true, nil
}

//--------------------------------------------
//Delete Contact
//---------------------------------------------
func (c *Contact) DeleteContact(db *gorm.DB, Id uint) (bool, error) {
	contact := en.TableContact{}
	db.First(&contact, Id)
	if contact.ID == 0 {
		return false, errors.New("contact id cannot be found")
	}
	db.Delete(&contact)
	return true, nil
}

//--------------------------------------------
//Get Contact By Id
//---------------------------------------------
func (c *Contact) ContactById(db *gorm.DB, Id uint) (bu.ContactBO, error) {

	contact := en.TableContact{}
	db.First(&contact, Id)
	return bu.ContactBO{Id: contact.ID, ContactTypeId: contact.ContactTypeId, Contact: contact.Contact, UpdatedAt: contact.UpdatedAt}, nil
}

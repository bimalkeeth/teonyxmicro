package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type IContactTypes interface {
	CreateContactType(db *gorm.DB, contactType bu.ContactTypeBO) (bool, error)
	UpdateContactType(db *gorm.DB, contactType bu.ContactTypeBO) (bool, error)
	DeleteContactType(db *gorm.DB, id uint) (bool, error)
	GetContactTypeById(db *gorm.DB, id uint) (bu.ContactTypeBO, error)
	GetContactTypeByName(db *gorm.DB, name string) (bu.ContactTypeBO, error)
	GetAll(db *gorm.DB) ([]bu.ContactTypeBO, error)
	GetAllNames(db *gorm.DB, namePart string) ([]bu.ContactTypeBO, error)
}

type ContactType struct{}

//-------------------------------------------
//Create instance to through above interface
//-------------------------------------------
func NewContactType() IContactTypes {
	return &ContactType{}
}

//-------------------------------------------
// Create Address type
//-------------------------------------------
func (c *ContactType) CreateContactType(db *gorm.DB, contactType bu.ContactTypeBO) (bool, error) {

	db.Create(&ent.TableContactType{ContactType: contactType.ContactType})
	return true, nil
}

//-------------------------------------------
//Update Contact Type
//-------------------------------------------
func (c *ContactType) UpdateContactType(db *gorm.DB, contactType bu.ContactTypeBO) (bool, error) {

	contacttype := ent.TableContactType{}
	db.First(&contacttype, contactType.Id)
	if contacttype.ID == 0 {
		return false, errors.New("contact type not found")
	}
	contacttype.ContactType = contactType.ContactType
	db.Save(&contactType)
	return true, nil
}

//-------------------------------------------
// Delete Contact Type
//-------------------------------------------
func (c *ContactType) DeleteContactType(db *gorm.DB, id uint) (bool, error) {
	found := ent.TableContactType{}
	db.First(&found, id)
	if found.ID == 0 {
		return false, errors.New("contact type not found")
	}
	db.Delete(&found)
	return true, nil

}

//-------------------------------------------
// Get Contact type by Id
//-------------------------------------------
func (c *ContactType) GetContactTypeById(db *gorm.DB, id uint) (bu.ContactTypeBO, error) {

	contactTypes := &ent.TableContactType{}
	db.First(&contactTypes, id)

	result := bu.ContactTypeBO{}
	if contactTypes.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.ContactTypeBO{ContactType: contactTypes.ContactType, Id: contactTypes.ID}, nil
}

//-------------------------------------------
// Get Contact Type by Name
//-------------------------------------------
func (c *ContactType) GetContactTypeByName(db *gorm.DB, name string) (bu.ContactTypeBO, error) {

	contactType := ent.TableContactType{}
	db.Where(&ent.TableContactType{ContactType: name}).First(&contactType)
	if contactType.ID == 0 {
		return bu.ContactTypeBO{}, errors.New("record not found")
	}
	return bu.ContactTypeBO{ContactType: contactType.ContactType, Id: contactType.ID}, nil
}

//-------------------------------------------
// Get All Contact type Name Value
//-------------------------------------------
func (c *ContactType) GetAll(db *gorm.DB) ([]bu.ContactTypeBO, error) {

	var contactTypes []ent.TableContactType
	var result []bu.ContactTypeBO

	db.Find(&contactTypes)
	for _, item := range contactTypes {
		result = append(result, bu.ContactTypeBO{ContactType: item.ContactType, Id: item.ID})
	}
	return result, nil

}

//-------------------------------------------
//Get all Names by name like
//-------------------------------------------
func (c *ContactType) GetAllNames(db *gorm.DB, namePart string) ([]bu.ContactTypeBO, error) {

	var contactTypes []ent.TableContactType
	db.Where("contacttype LIKE ?", "%"+namePart+"%").Find(&contactTypes)
	var result []bu.ContactTypeBO
	for _, item := range contactTypes {
		result = append(result, bu.ContactTypeBO{ContactType: item.ContactType, Id: item.ID})
	}
	return result, nil

}

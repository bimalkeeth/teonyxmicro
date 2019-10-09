package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type IOperatorContact interface {
	CreateOperatorContact(contactId uint, operatorId uint) (uint, error)
	UpdateOperatorContact(id uint, contactId uint, operatorId uint) (bool, error)
	DeleteOperatorContact(id uint) (bool, error)
	GetAllContactsByOperator(operatorId uint) ([]bu.OperatorContactsBO, error)
}

type OperatorContact struct {
	Db *gorm.DB
}

func NewOperatorContact(db *gorm.DB) OperatorContact {
	return OperatorContact{Db: db}
}

//-----------------------------------------------------
//Create operator contact
//-----------------------------------------------------
func (o *OperatorContact) CreateOperatorContact(contactId uint, operatorId uint) (uint, error) {
	opContact := ent.TableOperatorContacts{ContactId: contactId, OperatorId: operatorId}
	o.Db.Create(&opContact)
	return opContact.ID, nil
}

//-----------------------------------------------------
//Update operator contact
//-----------------------------------------------------
func (o *OperatorContact) UpdateOperatorContact(id uint, contactId uint, operatorId uint) (bool, error) {

	opContact := ent.TableOperatorContacts{}
	o.Db.First(&opContact, id)
	if opContact.ID == 0 {
		return false, errors.New("operator contact not found")
	}
	opContact.OperatorId = operatorId
	opContact.ContactId = contactId
	o.Db.Save(&opContact)
	return true, nil
}

//-----------------------------------------------------
//Delete operator contact
//-----------------------------------------------------
func (o *OperatorContact) DeleteOperatorContact(id uint) (bool, error) {
	opContact := ent.TableOperatorContacts{}
	o.Db.First(&opContact, id)
	if opContact.ID == 0 {
		return false, errors.New("operator contact not found")
	}
	o.Db.Delete(&opContact)
	return true, nil
}

//-----------------------------------------------------
//get contacts for operator
//-----------------------------------------------------
func (o *OperatorContact) GetAllContactsByOperator(operatorId uint) ([]bu.OperatorContactsBO, error) {

	var operators []ent.TableOperatorContacts
	var oprResults []bu.OperatorContactsBO

	o.Db.Preload("Contact").Preload("Operator").
		Where(&ent.TableOperatorContacts{OperatorId: operatorId}).Find(&operators)

	for _, item := range operators {

		oprResults = append(oprResults, bu.OperatorContactsBO{
			Id:         item.ID,
			ContactId:  item.ContactId,
			OperatorId: item.OperatorId,
			Contact: bu.ContactBO{
				Id:            item.ContactId,
				ContactTypeId: item.Contact.ContactTypeId,
				UpdatedAt:     item.Contact.UpdatedAt,
				Contact:       item.Contact.Contact,
			},
		})
	}
}

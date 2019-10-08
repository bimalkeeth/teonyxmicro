package business

import (
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

func (o *OperatorContact) UpdateOperatorContact(id uint, contactId uint, operatorId uint) (bool, error) {

}

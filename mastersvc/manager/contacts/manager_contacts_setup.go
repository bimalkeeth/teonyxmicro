package contacts

import (
	"log"
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
	ma "teonyxmicro/mastersvc/manager"
)

type IContactManager interface {
	CreateContact(con bu.ContactBO) (uint, error)
	UpdateContact(con bu.ContactBO) (bool, error)
	DeleteContact(id uint) (bool, error)
	ContactById(Id uint) (bu.ContactBO, error)
	CreateAddress(add bu.AddressBO) (uint, error)
	UpdateAddress(add bu.AddressBO) (bool, error)
	DeleteAddress(id uint) (bool, error)
	GetAddressById(id uint) (bu.AddressBO, error)
	GetAddressByName(name string) ([]bu.AddressBO, error)
}

type ContactManager struct{}

func New() IContactManager {
	return &ContactManager{}
}

var conFactory *bs.RuleFactory

func init() {
	conn, err := ma.Setconn()
	if err != nil {
		log.Fatal("error in contact manager initialisation")
	}
	conFactory = &bs.RuleFactory{Conn: conn}
}

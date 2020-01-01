package contacts

import (
	"log"
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
	ma "teonyxmicro/mastersvc/manager"
)

type IContactManager interface {
	CreateContact(con bu.ContactBO) (bool, uint)
	UpdateContact(con bu.ContactBO) bool
	DeleteContact(id uint) bool
	ContactById(Id uint) (bu.ContactBO, error)
	CreateAddress(add bu.AddressBO) (bool, uint)
	UpdateAddress(add bu.AddressBO) bool
	DeleteAddress(id uint) bool
	GetAddressById(id uint) (bu.AddressBO, error)
	GetAddressByName(name string) ([]bu.AddressBO, error)
}

type ContactManager struct{}

func New() *ContactManager {
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

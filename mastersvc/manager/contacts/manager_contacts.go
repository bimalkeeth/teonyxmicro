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

func NewContactManager() *ContactManager {
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

//---------------------------------------
//Create contact
//---------------------------------------
func (c *ContactManager) CreateContact(con bu.ContactBO) (bool, uint) {

	conFactory.Conn.Begin()
	contact := conFactory.New(bs.CContact).(bs.Contact)
	id, err := contact.CreateContact(con)
	if err != nil {
		conFactory.Conn.Rollback()
		return false, 0
	}
	conFactory.Conn.Commit()
	return true, id
}

//--------------------------------------
//Update Contact
//--------------------------------------
func (c *ContactManager) UpdateContact(con bu.ContactBO) bool {

	conFactory.Conn.Begin()
	contact := conFactory.New(bs.CContact).(bs.Contact)
	result, err := contact.UpdateContact(con)
	if err != nil {
		conFactory.Conn.Rollback()
		return false
	}
	conFactory.Conn.Commit()
	return result
}

//------------------------------------
//Delete Contact
//------------------------------------
func (c *ContactManager) DeleteContact(id uint) bool {

	conFactory.Conn.Begin()
	contact := conFactory.New(bs.CContact).(bs.Contact)
	result, err := contact.DeleteContact(id)
	if err != nil {
		conFactory.Conn.Rollback()
		return false
	}
	conFactory.Conn.Commit()
	return result
}

//------------------------------------
//Get Contact by Id
//------------------------------------
func (c *ContactManager) ContactById(Id uint) (bu.ContactBO, error) {

	contact := conFactory.New(bs.CContact).(bs.Contact)
	co, err := contact.ContactById(Id)
	if err != nil {

		return bu.ContactBO{}, err
	}
	return co, nil
}

//------------------------------------
//Create Address
//------------------------------------
func (c *ContactManager) CreateAddress(add bu.AddressBO) (bool, uint) {

	conFactory.Conn.Begin()
	address := conFactory.New(bs.CAddress).(bs.Address)
	res, err := address.CreateAddress(add)
	if err != nil {
		conFactory.Conn.Rollback()
		return false, 0
	}
	conFactory.Conn.Commit()
	return true, res
}

//------------------------------------
//Update Address
//------------------------------------
func (c *ContactManager) UpdateAddress(add bu.AddressBO) bool {

	conFactory.Conn.Begin()
	address := conFactory.New(bs.CAddress).(bs.Address)
	res, err := address.UpdateAddress(add)
	if err != nil {
		conFactory.Conn.Rollback()
		return false
	}
	conFactory.Conn.Commit()
	return res
}

//------------------------------------
//Delete Address
//------------------------------------
func (c *ContactManager) DeleteAddress(id uint) bool {

	conFactory.Conn.Begin()
	address := conFactory.New(bs.CAddress).(bs.Address)
	res, err := address.DeleteAddress(id)
	if err != nil {
		conFactory.Conn.Rollback()
		return false
	}
	conFactory.Conn.Commit()
	return res
}

//------------------------------------
//Get Address by AddressId
//------------------------------------
func (c *ContactManager) GetAddressById(id uint) (bu.AddressBO, error) {

	address := conFactory.New(bs.CAddress).(bs.Address)
	result, err := address.GetAddressById(id)
	if err != nil {
		return bu.AddressBO{}, err
	}
	return result, nil
}

//-------------------------------------
//Get Address By Name
//-------------------------------------
func (c *ContactManager) GetAddressByName(name string) ([]bu.AddressBO, error) {
	address := conFactory.New(bs.CAddress).(bs.Address)
	result, err := address.GetAddressByName(name)
	if err != nil {
		return result, err
	}
	return result, nil
}

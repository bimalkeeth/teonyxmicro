package manager

import (
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
	cn "teonyxmicro/mastersvc/connection"
)

type IContactManager interface {
	CreateContact(con bu.ContactBO) (bool, uint)
	UpdateContact(con bu.ContactBO) bool
	DeleteContact(id uint) bool
	ContactById(Id uint) (bu.ContactBO, error)
}

type ContactManager struct{}

//---------------------------------------
//Set database connection
//---------------------------------------
func setconn() (*gorm.DB, error) {
	conn := cn.New()
	db, err := conn.Open()
	if err != nil {
		return db, err
	}
	return db, nil
}

//---------------------------------------
//Create contact
//---------------------------------------
func (c *ContactManager) CreateContact(con bu.ContactBO) (bool, uint) {
	conn, err := setconn()
	if err != nil {
		return false, 0
	}
	conn.Begin()
	contact := bs.NewContact(conn)
	id, err := contact.CreateContact(con)
	if err != nil {
		conn.Rollback()
		return false, 0
	}
	conn.Commit()
	return true, id
}

//--------------------------------------
//Update Contact
//--------------------------------------
func (c *ContactManager) UpdateContact(con bu.ContactBO) bool {
	conn, err := setconn()
	if err != nil {
		return false
	}
	conn.Begin()
	contact := bs.NewContact(conn)
	result, err := contact.UpdateContact(con)
	if err != nil {
		conn.Rollback()
		return false
	}
	conn.Commit()
	return result
}

//------------------------------------
//Delete Contact
//------------------------------------
func (c *ContactManager) DeleteContact(id uint) bool {
	conn, err := setconn()
	if err != nil {
		return false
	}
	conn.Begin()
	contact := bs.NewContact(conn)
	result, err := contact.DeleteContact(id)
	if err != nil {
		conn.Rollback()
		return false
	}
	conn.Commit()
	return result
}

//------------------------------------
//Get Contact by Id
//------------------------------------
func (c *ContactManager) ContactById(Id uint) (bu.ContactBO, error) {
	conn, err := setconn()
	if err != nil {
		return bu.ContactBO{}, err
	}
	conn.Begin()
	contact := bs.NewContact(conn)
	co, err := contact.ContactById(Id)
	if err != nil {
		return bu.ContactBO{}, err
	}
	return co, nil
}

package fleets

import (
	"log"
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
	ma "teonyxmicro/mastersvc/manager"
)

var flFac *bs.RuleFactory

func init() {
	conn, err := ma.Setconn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	flFac = &bs.RuleFactory{Conn: conn}
}

type IFleetManager interface {
	CreateFleet(bo bu.FleetBO) (bu.FleetBO, error)
	UpdateFleet(bo bu.FleetBO) (bool, error)
	DeleteFleet(id uint) (bool, error)
	GetFleetById(id uint) (bu.FleetBO, error)
	CreateFleetContact(fleetId uint, contactId uint) (uint, error)
	UpdateFleetContact(id uint, fleetId uint, contactId uint) (bool, error)
	DeleteFleetContact(id uint) (bool, error)
	GetContactByFleetId(fleetId uint) ([]bu.FleetContactBO, error)
	CreateFleetLocation(fleetId uint, addressId uint) (uint, error)
	UpdateFleetLocation(id uint, fleetId uint, addressId uint) (bool, error)
	DeleteFleetLocation(id uint) (bool, error)
	GetLocationByFleetId(fleetId uint) ([]bu.FleetAddressBO, error)
}

type FleetManager struct{}

func New() *FleetManager {
	return &FleetManager{}
}

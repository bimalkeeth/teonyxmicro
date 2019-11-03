package masters

import (
	"log"
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
	ma "teonyxmicro/mastersvc/manager"
)

var masFac *bs.RuleFactory

func init() {
	conn, err := ma.Setconn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	masFac = &bs.RuleFactory{Conn: conn}
}

type IMasterManager interface {
	CreateCompany(company bu.CompanyBO) (uint, error)
	UpdateCompany(company bu.CompanyBO) (bool, error)
	DeleteCompany(id uint) (bool, error)
	CreateAddressType(addressType bu.AddressTypeBO) (uint, error)
	UpdateAddressType(addressType bu.AddressTypeBO) (bool, error)
	DeleteAddressType(id uint) (bool, error)
	GetAddressTypeById(id uint) (bu.AddressTypeBO, error)
	GetAddressTypeByName(name string) (bu.AddressTypeBO, error)
	GetAllAddressTypes() ([]bu.AddressTypeBO, error)
	GetAllAddressTypeNames(namePart string) ([]bu.AddressTypeBO, error)
}

type MasterManager struct{}

func NewMasterManager() *MasterManager {
	return &MasterManager{}
}

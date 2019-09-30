package business

import (
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

//----------------------------------------------
//interface for company
//----------------------------------------------
type ICompany interface {
	CreateCompany(db *gorm.DB, company bu.CompanyBO) (bool, error)
	UpdateCompany(db *gorm.DB, company bu.CompanyBO) (bool, error)
	DeleteCompany(db *gorm.DB, id uint) (bool, error)
}

type Company struct{}

func NewCompany() ICompany { return Company{} }

//---------------------------------------------
//Create Company
//---------------------------------------------
func (c Company) CreateCompany(db *gorm.DB, company bu.CompanyBO) (bool, error) {

	db.Create(&ent.TableCompany{Name: company.Name,
		AddressId:  company.AddressId,
		ContractId: company.ContactId})
	return true, nil
}

func (c Company) UpdateCompany(db *gorm.DB, company bu.CompanyBO) (bool, error) {

	return true, nil
}

func (c Company) DeleteCompany(db *gorm.DB, id uint) (bool, error) {

	return true, nil
}

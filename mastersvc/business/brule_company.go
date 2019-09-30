package business

import (
	"errors"
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

func NewCompany() *Company { return &Company{} }

//----------------------------------------------
//Create Company
//----------------------------------------------
func (c Company) CreateCompany(db *gorm.DB, company bu.CompanyBO) (bool, error) {

	db.Create(&ent.TableCompany{Name: company.Name,
		AddressId:  company.AddressId,
		ContractId: company.ContactId})
	return true, nil
}

//-----------------------------------------------
//Update company
//-----------------------------------------------
func (c Company) UpdateCompany(db *gorm.DB, company bu.CompanyBO) (bool, error) {

	com := &ent.TableCompany{}
	db.First(com, company.Id)
	if com.ID == 0 {
		return false, errors.New("company can not be found")
	}
	com.ContractId = company.ContactId
	com.AddressId = company.AddressId
	com.Name = company.Name
	db.Save(com)
	return true, nil
}

//-----------------------------------------------
//Delete company
//-----------------------------------------------
func (c Company) DeleteCompany(db *gorm.DB, id uint) (bool, error) {

	com := ent.TableCompany{}
	db.First(&com, id)
	if com.ID == 0 {
		return false, errors.New("company type not found")
	}
	db.Delete(&com)
	return true, nil
}

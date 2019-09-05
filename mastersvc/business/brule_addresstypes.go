package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type AddressTypes interface {
	CreateAddressType(db *gorm.DB, addressType bu.AddressTypeBO) (bool, error)
	UpdateAddressType(db *gorm.DB, addressType bu.AddressTypeBO) (bool, error)
	DeleteAddressType(db *gorm.DB, id uint) (bool, error)
	GetAddressTypeById(db *gorm.DB, id uint) (bu.AddressTypeBO, error)
	GetAddressTypeByName(db *gorm.DB, name string) (bu.AddressTypeBO, error)
	GetAll(db *gorm.DB) ([]bu.AddressTypeBO, error)
	GetAllNemes(db *gorm.DB, namePart string) ([]bu.AddressTypeBO, error)
}
type AddressType struct{}

func New() *AddressType { return &AddressType{} }

func (at *AddressType) CreateAddressType(db *gorm.DB, addressType bu.AddressTypeBO) (bool, error) {

	result := db.NewRecord(ent.TableAddressType{AddressType: addressType.Name})
	if !result {
		return result, errors.New("address type creation not successful")
	}
	return result, nil
}

func (at *AddressType) UpdateAddressType(db *gorm.DB, addressType bu.AddressTypeBO) (bool, error) {

	addressTypes := &ent.TableAddressType{}
	db.First(&addressTypes, addressType.Id)
	if addressTypes.ID == 0 {
		return false, nil
	}
	addressTypes.Name = addressType.Name
	db.Save(&addressTypes)
	return true, nil
}

func (at *AddressType) DeleteAddressType(db *gorm.DB, id uint) (bool, error) {

	addressTypes := &ent.TableAddressType{}
	db.First(&addressTypes, id)

	if addressTypes.ID == 0 {
		return false, errors.New("the record not exists in the storage")
	}
	db.Delete(&addressTypes)
	return true, nil
}

func (at *AddressType) GetAddressTypeById(db *gorm.DB, id uint) (bu.AddressTypeBO, error) {

	addressTypes := &ent.TableAddressType{}
	db.First(&addressTypes, id)

	result := bu.AddressTypeBO{}
	if addressTypes.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.AddressTypeBO{Name: addressTypes.Name, Id: addressTypes.ID}, nil
}

func (at *AddressType) GetAddressTypeByName(db *gorm.DB, name string) (bu.AddressTypeBO, error) {

	addressType := ent.TableAddressType{}
	db.Where(&ent.TableAddressType{Name: name}).First(&addressType)
	if addressType.ID == 0 {
		return bu.AddressTypeBO{}, errors.New("record not found")
	}
	return bu.AddressTypeBO{Name: addressType.Name, Id: addressType.ID}, nil
}

func (at *AddressType) GetAll(db *gorm.DB) ([]bu.AddressTypeBO, error) {

	var addressTypes []ent.TableAddressType
	var result []bu.AddressTypeBO

	db.Find(&addressTypes)
	for _, item := range addressTypes {
		result = append(result, bu.AddressTypeBO{Name: item.Name, Id: item.ID})
	}
	return result, nil
}

func (at *AddressType) GetAllNemes(db *gorm.DB, namePart string) ([]bu.AddressTypeBO, error) {
	var addressTypes []ent.TableAddressType
	db.Where("name LIKE ?", "%"+namePart+"%").Find(&addressTypes)
	var result []bu.AddressTypeBO
	for _, item := range addressTypes {
		result = append(result, bu.AddressTypeBO{Name: item.Name, Id: item.ID})
	}
	return result, nil

}

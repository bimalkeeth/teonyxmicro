package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "teonyxmicro/mastersvc/bucontracts"
	ent "teonyxmicro/mastersvc/entities"
)

type IAddress interface {
	CreateAddress(db *gorm.DB, address bu.AddressBO) (bool, error)
	UpdateAddress(db *gorm.DB, address bu.AddressBO) (bool, error)
	DeleteAddress(db *gorm.DB, id uint) (bool, error)
	GetAddressById(db *gorm.DB, id uint) (bu.AddressBO, error)
	GetAddressByName(db *gorm.DB, name string) ([]bu.AddressBO, error)
}

type Address struct{}

//---------------------------------------------------
//Create address
//---------------------------------------------------
func (a *Address) CreateAddress(db *gorm.DB, address bu.AddressBO) (bool, error) {

	result := db.NewRecord(&ent.TableAddress{CountryId: address.CountryId,
		AddressTypeId: address.AddressTypeId,
		StateId:       address.StateId,
		Location:      address.Location,
		Address:       address.Address,
		Street:        address.Street,
		Suburb:        address.Suburb})

	if !result {
		return false, errors.New("error in creating address")
	}
	return true, nil
}

//---------------------------------------------------
//Update Address
//---------------------------------------------------
func (a *Address) UpdateAddress(db *gorm.DB, address bu.AddressBO) (bool, error) {

	addr := &ent.TableAddress{}
	db.First(addr, address.Id)
	if addr.ID == 0 {
		return false, errors.New("address not found")
	}
	addr.Suburb = address.Suburb
	addr.Street = address.Street
	addr.Address = address.Address
	addr.Location = address.Location
	addr.StateId = address.StateId
	addr.AddressTypeId = address.AddressTypeId
	addr.CountryId = address.CountryId
	db.Save(addr)
	return true, nil
}

//---------------------------------------------------
//Delete Address
//---------------------------------------------------
func (a *Address) DeleteAddress(db *gorm.DB, id uint) (bool, error) {

	address := &ent.TableAddress{}
	db.First(&address, id)

	if address.ID == 0 {
		return false, errors.New("the record not exists in the storage")
	}
	db.Delete(&address)
	return true, nil
}

//----------------------------------------------------
//Get Address by Id
//----------------------------------------------------
func (a *Address) GetAddressById(db *gorm.DB, id uint) (bu.AddressBO, error) {

	address := &ent.TableAddress{}
	db.First(&address, id)

	result := bu.AddressBO{}
	if address.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.AddressBO{CountryId: address.CountryId,
		AddressTypeId: address.AddressTypeId,
		StateId:       address.StateId,
		Location:      address.Location,
		Address:       address.Address,
		Street:        address.Street,
		Suburb:        address.Suburb,
		Id:            address.ID,
	}, nil
}

//----------------------------------------------------
//Get Address by Name
//----------------------------------------------------
func (a *Address) GetAddressByName(db *gorm.DB, name string) ([]bu.AddressBO, error) {

	var address []ent.TableAddress
	db.Where("name LIKE ?", "%"+name+"%").Find(&address)
	var result []bu.AddressBO
	for _, item := range address {
		result = append(result, bu.AddressBO{CountryId: item.CountryId,
			AddressTypeId: item.AddressTypeId,
			StateId:       item.StateId,
			Location:      item.Location,
			Address:       item.Address,
			Street:        item.Street,
			Suburb:        item.Suburb,
			Id:            item.ID,
		})
	}
	return result, nil
}

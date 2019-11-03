package masters

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

//----------------------------------------
//Create Company
//----------------------------------------
func (m *MasterManager) CreateCompany(company bu.CompanyBO) (uint, error) {

	master := masFac.New(bs.CCompany).(bs.Company)
	masFac.Conn.Begin()
	res, err := master.CreateCompany(company)
	if err != nil {
		masFac.Conn.Rollback()
		return 0, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Update Company
//-----------------------------------------
func (m *MasterManager) UpdateCompany(company bu.CompanyBO) (bool, error) {
	master := masFac.New(bs.CCompany).(bs.Company)
	masFac.Conn.Begin()
	res, err := master.UpdateCompany(company)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Delete Company
//-----------------------------------------
func (m *MasterManager) DeleteCompany(id uint) (bool, error) {
	master := masFac.New(bs.CCompany).(bs.Company)
	masFac.Conn.Begin()
	res, err := master.DeleteCompany(id)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Create Address Type
//-----------------------------------------
func (m *MasterManager) CreateAddressType(addressType bu.AddressTypeBO) (uint, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	masFac.Conn.Begin()
	res, err := master.CreateAddressType(addressType)
	if err != nil {
		masFac.Conn.Rollback()
		return 0, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Update Address Type
//-----------------------------------------
func (m *MasterManager) UpdateAddressType(addressType bu.AddressTypeBO) (bool, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	masFac.Conn.Begin()
	res, err := master.UpdateAddressType(addressType)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//----------------------------------------
//Delete Address Type
//----------------------------------------
func (m *MasterManager) DeleteAddressType(id uint) (bool, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	masFac.Conn.Begin()
	res, err := master.DeleteAddressType(id)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Get AddressType by Id
//-----------------------------------------
func (m *MasterManager) GetAddressTypeById(id uint) (bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAddressTypeById(id)
	if err != nil {
		return bu.AddressTypeBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Address By Name
//-----------------------------------------
func (m *MasterManager) GetAddressTypeByName(name string) (bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAddressTypeByName(name)
	if err != nil {
		return bu.AddressTypeBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Address
//-----------------------------------------
func (m *MasterManager) GetAllAddressTypes() ([]bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAll()
	if err != nil {
		return []bu.AddressTypeBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Address By Name like
//-----------------------------------------
func (m *MasterManager) GetAllAddressTypeNames(namePart string) ([]bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAllNames(namePart)
	if err != nil {
		return []bu.AddressTypeBO{}, err
	}
	return res, nil
}
